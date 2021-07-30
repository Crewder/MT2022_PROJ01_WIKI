package helpers

import (
	"log"
)

func HandleError(status int, err error) {
	if err != nil {
		log.Fatal(status, err)
	}
}
