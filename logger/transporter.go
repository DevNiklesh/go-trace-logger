package logger

type Transporter interface {
	Transport(level Level, message string) error
}
