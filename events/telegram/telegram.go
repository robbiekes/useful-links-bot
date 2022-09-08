package telegram

import "github.com/robbiekes/LinksBot/client/telegram"

type EventProcessor struct {
	tgClient *telegram.Client
	offset   int
}

func NewEventProcessor(tgClient *telegram.Client) *EventProcessor {
	return &EventProcessor{tgClient: tgClient}
}
