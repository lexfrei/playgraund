package main

import (
	"github.com/lexfrei/playgraund/api"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Create an instance of your API implementation
	myApi := api.IBolitAPI{}

	// Register the handlers
	api.RegisterHandlers(e, &myApi)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
