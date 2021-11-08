package domain

type Log struct {
	Id      string
	Content string
}

type LogRepository interface {
	Store(log *Log) error
}
