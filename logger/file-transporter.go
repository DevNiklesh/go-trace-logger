package logger

import (
	"os"
	"sync"
)

type FileTransporter struct {
	mu   sync.Mutex
	file *os.File
}

func NewFileTransporter(filePath string) (*FileTransporter, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return &FileTransporter{file: file}, nil
}

func (t *FileTransporter) Transport(level Level, message string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	_, err := t.file.WriteString(message + "\n")
	if err != nil {
		return err
	}
	return nil
}

func (t *FileTransporter) Close() error {
	return t.file.Close()
}
