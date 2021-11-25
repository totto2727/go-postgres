package main

import (
	"github.com/labstack/echo"
	"go_api/controller"
)

func main() {
	e := echo.New()
	e.GET("/users", controller.GetUser)
	e.Logger.Fatal(e.Start(":8080"))
}
