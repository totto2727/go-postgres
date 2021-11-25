package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name string `query:"name" json:"name"`
	Age  int    `query:"age" json:"age"`
}

func GetUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
