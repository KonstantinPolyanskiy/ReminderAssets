package events

const (
	Unknown Type = iota
	Message
)

type Type int

type Fetcher interface {
	Fetch(limit int)
}
type Processor interface {
	Process(e Event) error
}
type Event struct {
	Type Type
	Text string
}
