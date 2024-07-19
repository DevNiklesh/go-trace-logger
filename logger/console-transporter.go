package logger

import "fmt"

type ConsoleTransporter struct{}

func (t *ConsoleTransporter) Transport(level Level, message string) error {
	fmt.Println(message)
	return nil
}
