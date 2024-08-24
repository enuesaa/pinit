package repository

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(0)
}

type LogRepositoryInterface interface {
	Info(message string, a...any)
	Errf(err error, message string, a...any)
	Err(err error)
}

type LogRepository struct {}

func (l *LogRepository) Info(message string, a...any) {
	text := fmt.Sprintf(message, a...)
	log.Println(text)
}

func (l *LogRepository) Errf(err error, message string, a...any) {
	text := fmt.Sprintf("Error: %s because `%s`", message, err.Error())
	// log.Panicln を呼ぶと error message が2回表示されるので panic を直接呼ぶ
	panic(text)
}

func (l *LogRepository) Err(err error) {
	text := fmt.Sprintf("Error: %s", err.Error())
	panic(text)	
}
