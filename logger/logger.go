package logger

import "fmt"

type Logger struct {
	Level          Level
	Transporter    Transporter
	Formatter      Formatter
	DefaultMessage string
}

func NewLogger(level Level, transporter Transporter, formatter Formatter, defaultMessage string) *Logger {
	return &Logger{
		level,
		transporter,
		formatter,
		defaultMessage,
	}
}

func (l *Logger) WithDefaultMessage(message string) *Logger {
	l.DefaultMessage = l.DefaultMessage + " " + message
	return NewLogger(l.Level, l.Transporter, l.Formatter, l.DefaultMessage)
}

func (l *Logger) log(level Level, message string) {
	if level < l.Level {
		return
	}

	msg, err := l.Formatter.Format(level, fmt.Sprintf("%s %s", l.DefaultMessage, message))
	if err != nil {
		return
	}

	err = l.Transporter.Transport(level, msg)
	if err != nil {
		return
	}
}

func (l *Logger) Debug(message string) {
	l.log(Debug, message)
}
func (l *Logger) Info(message string) {
	l.log(Info, message)
}
func (l *Logger) Warn(message string) {
	l.log(Warning, message)
}
func (l *Logger) Error(message string) {
	l.log(Error, message)
}
func (l *Logger) Fatal(message string) {
	l.log(Fatal, message)
}
