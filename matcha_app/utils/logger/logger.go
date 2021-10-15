package logger

import (
	"log"
	"os"
)

var globalLogger *log.Logger

func init() {
	globalLogger = log.New(os.Stdout, "Global:\t", log.LstdFlags)
}

// Get global logger singleton instance
func Get() *log.Logger {
	return globalLogger
}
