package log

import (
	"log"
	"os"
)

const (
	file = "log.txt"
)

var Logger *log.Logger

func init() {
	f, ferr := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if ferr != nil {
		log.Fatal(ferr)
	}
	Logger = log.New(f, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
}
