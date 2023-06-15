package services

import "fmt"

type Logger struct{}

func ProvideLogger() Logger {
	return Logger{}
}

func (l *Logger) Logged() {
	fmt.Println("logger was called")
}
