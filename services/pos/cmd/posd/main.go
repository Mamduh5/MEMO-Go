package main

import (
	"log"

	"memo-go/services/pos/internal/app"
)

func main() {
	if err := app.NewApp(); err != nil {
		log.Fatal(err)
	}
}
