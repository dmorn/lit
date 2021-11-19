package log

import (
	"fmt"
	"log"
	"os"
)

func Errore(err error) {
	log.Printf("ERROR * %v", err)
}

func Errorf(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	Errore(err)
}

func Fatale(err error) {
	Errore(err)
	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	Fatale(err)
}

func Info(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Warn(format string, args ...interface{}) {
	log.Printf("WARN * "+format, args...)
}

func Default() *log.Logger { return log.Default() }
