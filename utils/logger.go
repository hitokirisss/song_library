package utils

import (
	"log"
)

func Debug(msg string) {
	log.Println("DEBUG:", msg)
}

func Info(msg string) {
	log.Println("INFO:", msg)
}
