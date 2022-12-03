package main

import "github.com/Subha-Research/koham/app"

func main() {
	app := app.SetupApp()
	app.Listen(":8080")
}
