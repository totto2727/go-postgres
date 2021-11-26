package controller

import (
	"go_api/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetUser(c echo.Context) error {
	query := new(model.User)
	if err := c.Bind(query); err != nil {
		return err
	}
	result := model.GetUsers(query)
	return c.JSON(http.StatusOK, result)
}

func PostUser(c echo.Context) error {
	query := new(model.User)
	if err := c.Bind(query); err != nil {
		return err
	}
	model.CreateUser(query)
	return c.JSON(http.StatusOK, "success")
}

func UpdateUsers(c echo.Context) error {
	query := new(model.User)
	if err := c.Bind(query); err != nil {
		return err
	}
	model.UpdateUsers(query)
	return c.JSON(http.StatusOK, "success")
}

func DeleteUser(c echo.Context) error {
	query := new(model.User)
	if err := c.Bind(query); err != nil {
		return err
	}
	model.DeleteUser(query)
	return c.NoContent(http.StatusNoContent)
}
