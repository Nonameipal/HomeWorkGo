package main

import "go-shop/internal/controller"

func main() {
	app := controller.NewController()
	app.Run()
}
