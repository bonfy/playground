package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World from echo")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	return c.String(http.StatusOK, fmt.Sprintf("Cat name: %s type: %s", catName, catType))
}

func main() {
	e := echo.New()

	e.GET("/", helloHandler)
	e.GET("/cats", getCats)

	e.Start(":8888")
}
