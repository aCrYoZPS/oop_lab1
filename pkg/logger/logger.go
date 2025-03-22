package logger

import (
	"log"
)

func Fatal(message string) {
	log.Fatalf("[FATAL] %s\n", message)
}

func Error(message string) {
	log.Printf("[ERROR] %s\n", message)
}

func Info(message string) {
	log.Printf("[INFO] %s\n", message)
}
