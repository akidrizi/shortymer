package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Welcome to Go URL Shortener with Redis !🚀",
		})
	})

	app.Logger.Fatal(app.Start(":1323"))
}
