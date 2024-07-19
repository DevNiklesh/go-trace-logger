package logger

type Formatter interface {
	Format(level Level, message string) (string, error)
}
