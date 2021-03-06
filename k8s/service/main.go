package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/health/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"status": "OK"}`)
	})
	e.Logger.Fatal(e.Start(":8000"))
}
