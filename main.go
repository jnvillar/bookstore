package main

import (
	"bookstore/app"
	"bookstore/config"
)

func main() {
	app := app.NewApp(config.DevConfig())
	app.Init()
}
