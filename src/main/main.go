package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func hello(c echo.Context) error{
	return c.String(http.StatusOK, "hello from the web side")
}

func getCats(c echo.Context) error{
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string"{
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is: %s\n and his type is :%s\n", catName, catType))
	}

	if dataType == "json"{
		return c.JSON(http.StatusOK, map[string]string{
			"name":catName,
			"type":catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error" : "you need to let us know if you want json or string data",
	})

}

func main(){
	fmt.Println("Welcome to the go server")

	// Create a new echo object
	e := echo.New()

	e.GET("/", hello)
	e.GET("/cats/:data", getCats)

	e.Start(":8000")
}