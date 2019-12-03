package logger

import "../config"

type Interface interface {
	Err(str string)
	Info(str string)
}

// --------------------------

func GetLogger() Interface {
	switch config.LoggerType {
	case "sentry":
		return &Sentry{}
	case "filesystem":
		return &FileSystem{}
	case "elk":
		return &ELK{}
	case "elastic":
		return &Elastic{}
	}

	panic("Попытка получить неизвестный логгер")
}

// --------------------------

type Sentry struct {
	Interface
}

func (object *Sentry) Err(error string) {
	// подключение к сентри по DSN, запись ошибки
}

func (object *Sentry) Info(error string) {
	// подключение к сентри по DSN, запись ошибки
}

// --------------------------

type FileSystem struct {
	Interface
}

func (object *FileSystem) Err(error string) {
	// открывает файл на диске, пишет туда ошибку, закрывает файл
}

func (object *FileSystem) Info(error string) {
	// открывает файл на диске, пишет туда ошибку, закрывает файл
}

// --------------------------

type ELK struct {
	Interface
}

func (object *ELK) Err(error string) {
	// записывает ошибку в elastic через logstash
}

func (object *ELK) Info(error string) {
	// записывает ошибку в elastic через logstash
}

// --------------------------

type Elastic struct {
	Interface
}

func (object *Elastic) Err(error string) {
	// записывает ошибку в elastic
}

func (object *Elastic) Info(error string) {
	// записывает ошибку в elastic
}
