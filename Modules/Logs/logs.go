package logutil

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func Init(logToFile bool) {
	var output io.Writer = os.Stdout
	if logToFile {
		file, err := os.OpenFile("resolver.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		output = io.MultiWriter(os.Stdout, file)
	}

	Logger = log.New(output, "[HopZero] ", log.LstdFlags)
}
