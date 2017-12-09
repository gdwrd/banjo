package main

import "banjo"

func main() {
	app = banjo.New(banjo.DefaultConfig())

	app.Get("/", func(req banjo.Request) {
		return app.JSON(banjo.M{"foo": "bar"})
	})

	app.Get("/about", func(req banjo.Request) {
		return app.HTML("<h1>Hello form Banjo</h1>")
	})

	app.Run()
}
