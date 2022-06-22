package main

import (
	"fmt"

	"github.com/datshiro/crud/internal/app"
)

func main() {
	app := app.NewApp()

	app.Parse()
	fmt.Println(app)
}
