package main

import (
	app "github.com/Subha-Research/svasthfamily-koham/app"
)

func main() {
	f_app := app.InitFiberApplication()
	app := app.KohamApp{}
	app.App = f_app

	fiber_app := app.SetupApp()
	fiber_app.Listen(":8080")
}
