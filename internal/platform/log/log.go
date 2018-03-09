package log

import (
	"fmt"
	"log"
	"os"
)

//StdOutLog prints error to stdout
func StdOutLog(err error) error {
	fmt.Println(err)
	return nil
}

//FileLog stores error in log file
func FileLog(err error) error {
	return nil
}

//MailLog sends email with error in body
func MailLog(err error) error {
	return nil
}

//Fatal calls logFunc with error followed by os.Exit(1)
func Fatal(logFunc func(err error) error, err error) {
	if err := logFunc(err); err != nil {
		log.Fatal(err)
	}
	os.Exit(1)
}
