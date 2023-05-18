package loggerutils

import (
	"log"
	"os"
)

func GetStdErrorLogger() *log.Logger {
	return log.New(os.Stderr, "", log.LstdFlags)
}

func GetStdOutputLogger() *log.Logger {
	return log.New(os.Stdout, "", 0)
}
