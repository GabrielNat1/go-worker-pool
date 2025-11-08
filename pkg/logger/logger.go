package logger

import (
	"fmt"
	"time"
)

func Info(msg string) {
	fmt.Printf("[INFO] %s | %s\n", time.Now().Format(time.RFC3339), msg)
}

func Error(msg string) {
	fmt.Printf("[ERROR] %s | %s\n", time.Now().Format(time.RFC3339), msg)
}

func Warn(msg string) {
	fmt.Printf("[WARN] %s | %s\n", time.Now().Format(time.RFC3339), msg)
}
