package telegram

import (
	"github.com/robbiekes/LinksBot/client/telegram"
	"github.com/robbiekes/LinksBot/events"
	"github.com/robbiekes/LinksBot/lib/e"
	"github.com/robbiekes/LinksBot/linksbot/repository"
)

type EventProcessor struct {
	tgClient *telegram.Client
	offset   int
	storage  repository.Storage
}

func NewEventProcessor(tgClient *telegram.Client, storage repository.Storage) *EventProcessor {
	return &EventProcessor{
		tgClient: tgClient,
		storage:  storage,
	}
}

func (p *EventProcessor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tgClient.GetUpdates(p.offset, limit)
	if err != nil {
		return nil, e.WrapErr("error when getting events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, upd := range updates {
		res = append(res, event(upd))
	}

	p.offset = updates[len(updates)-1].UpdateId + 1

	return res, nil

}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)

	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}

	if updType == events.Message {
		res.Meta = telegram.Meta{
			ChatId:   upd.Message.Chat.Id,
			Username: upd.Message.From.Username,
		}
	}

	return res
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text
}

func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}
	return events.Message
}
