package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	e := echo.New()
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	e.GET("/", func(c echo.Context) error {
		log.Println("Hello, World!")
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
