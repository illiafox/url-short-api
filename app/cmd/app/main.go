package main

import (
	apps "ozon-url-shortener/app/internal/app"
)

func main() {
	app := apps.Init()
	app.Run()
}
