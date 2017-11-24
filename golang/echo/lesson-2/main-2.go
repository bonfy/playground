package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// Cat Model
type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World from echo")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	outputType := c.Param("output")

	if outputType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Cat name: %s type: %s", catName, catType))
	}

	if outputType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "please give valid output type",
	})

}

func addCats(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body for addCats: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Failed unmarshaling in addCats: %s\n", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("this is your cat: %#v\n", cat)
	return c.String(http.StatusOK, "we got your cat!")
}

func main() {
	e := echo.New()

	e.GET("/", helloHandler)
	e.GET("/cats/:output", getCats)

	e.POST("/cats", addCats)

	e.Start(":8888")
}
