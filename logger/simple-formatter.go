package logger

import (
	"fmt"
	"time"
)

type SimpleFormatter struct{}

func (f *SimpleFormatter) Format(level Level, message string) (string, error) {
	str := fmt.Sprintf("%v [%s] %s", time.Now().Format(time.RFC3339), level.String(), message)
	return str, nil
}
