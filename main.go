package main

import app "github.com/Subha-Research/svasthfamily-koham/app"

func main() {
	app := app.KohamApp{}
	fiber_app := app.SetupApp()
	fiber_app.Listen(":8080")
}
