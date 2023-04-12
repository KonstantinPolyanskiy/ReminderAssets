package storage

import "time"

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) error
	Remove(p *Page) error
	ISExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
	Created  time.Time
}
