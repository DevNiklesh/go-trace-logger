package main

import (
	l "go-logger/logger"
	"log"
)

func main() {
	consoleTransporter := &l.ConsoleTransporter{}
	// consoleLogger := l.NewLogger(l.Info, &l.ConsoleTransporter{}, &l.SimpleFormatter{})
	// consoleLogger.Debug("This is Debug")
	// consoleLogger.Info("This is info")
	// consoleLogger.Warn("This is Warning")
	// consoleLogger.Error("This is Error")
	// consoleLogger.Fatal("This is Fatal")

	fileTransporter, err := l.NewFileTransporter("logs.txt")
	if err != nil {
		log.Fatalln("Unable to create logs.txt file", err)
	}

	// fileLogger := l.NewLogger(l.Warning, fileTransporter, &l.SimpleFormatter{})
	// fileLogger.Debug("This is Debug")
	// fileLogger.Info("This is info")
	// fileLogger.Warn("This is Warning")
	// fileLogger.Error("This is Error")
	// fileLogger.Fatal("This is Fatal")

	multiTransporter := l.NewMultiTransporter(consoleTransporter, fileTransporter)
	multiLogger := l.NewLogger(l.Warning, multiTransporter, &l.SimpleFormatter{}, "Parent")
	multiLogger.Debug("This is Debug")
	multiLogger.Info("This is info")
	multiLogger.Warn("This is Warning")
	multiLogger.Error("This is Error")
	multiLogger.Fatal("This is Fatal")

	multiChildLogger := multiLogger.WithDefaultMessage("Child")
	multiChildLogger.Debug("This is Debug")
	multiChildLogger.Info("This is info")
	multiChildLogger.Warn("This is Warning")
	multiChildLogger.Error("This is Error")
	multiChildLogger.Fatal("This is Fatal")

	grandChildLogger := multiChildLogger.WithDefaultMessage("GrandChild")
	grandChildLogger.Debug("This is Debug")
	grandChildLogger.Info("This is info")
	grandChildLogger.Warn("This is Warning")
	grandChildLogger.Error("This is Error")
	grandChildLogger.Fatal("This is Fatal")
}
