package helpers

import (
	"log"
)

func HandleError(status int, err error, bool bool) {
	if err != nil || bool != false {
		log.Fatal(status, err)
	}
}
