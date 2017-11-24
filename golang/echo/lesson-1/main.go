package main

import "github.com/labstack/echo"
import "net/http"

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World from echo")
	})

	e.Start(":8888")
}
