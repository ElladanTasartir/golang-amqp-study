package util

import "log"

func FailOnError(err error, message string) {
	if err != nil {
		log.Fatalf("Error %s: %s", message, err)
	}
}
