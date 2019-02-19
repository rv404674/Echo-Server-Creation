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

	return c.String(http.StatusOK, fmt.Sprintf("your cat name is: %s\n and his type is :%s\n", catName, catType))

}

func main(){
	fmt.Println("Welcome to the go server")

	// Create a new echo object
	e := echo.New()

	e.GET("/", hello)
	e.GET("/cats", getCats)

	e.Start(":8000")
}