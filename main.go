package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type response struct {
	Message string `json:"message"`
}

func main() {

	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {

		return c.JSON(http.StatusOK, response{Message: "hello"})
	})

	e.GET("/goodbye", func(c echo.Context) error {

		return c.JSON(http.StatusOK, response{Message: "goodbye"})
	})

	e.GET("/error", func(c echo.Context) error {

		return c.JSON(http.StatusBadRequest, response{Message: "you did something wrong"})
	})

	log.Fatal(e.Start(":8000"))
}
