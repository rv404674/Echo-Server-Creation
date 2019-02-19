package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func hello(c echo.Context) error{
	return c.String(http.StatusOK, "hello from the web side")
}

func main(){
	fmt.Println("Welcome to the go server")

	// Create a new echo object
	e := echo.New()

	e.GET("/", hello)

	e.Start(":8000")
}