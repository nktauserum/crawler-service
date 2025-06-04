package main

import (
	"log"

	"github.com/nktauserum/crawler-service/internal/app"
)

func main() {
	app := app.NewApplication(8090)

	err := app.Run()
	if err != nil {
		log.Fatalf("error starting the app: %s\n", err.Error())
	}
}
