package logger

import "log"

func Info(message string) {
	log.Println("[INFO]", message)
}

func Error(message string) {
	log.Println("[ERROR]", message)
}
