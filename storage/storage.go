package storage

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	isPresent(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}