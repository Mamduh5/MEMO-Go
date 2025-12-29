package main

import (
	"log"

	"memo-go/services/auth/internal/app"
)

func main() {
	if err := app.NewApp(); err != nil {
		log.Fatal(err)
	}
}
