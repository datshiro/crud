package main

import (
	"log"

	"github.com/datshiro/crud/internal/app"
)

func main() {
	app := app.NewApp()

	app.Parse()

	app.ConfigMiddleware()

	app.ConfigLogLevel()

	app.RegisterHandlers()

	// run app
	if err := app.Run(); err != nil {
		log.Fatalf("fail to start, err=%v", err)
	}
}
