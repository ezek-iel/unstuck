package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func reportServerError(message string, err error, c *echo.Context) error {
	c.Logger().Error(message, "error", err)
	return c.JSON(http.StatusInternalServerError, struct{ message string }{message: "An error occured on our end"})
}

func reportClientError(message map[string]string, c *echo.Context) error {
	c.Logger().Error("client error")
	return c.JSON(http.StatusBadRequest, message)
}