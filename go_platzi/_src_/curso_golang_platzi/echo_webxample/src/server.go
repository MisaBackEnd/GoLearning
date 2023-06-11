package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	var a int
	fmt.Printf("%d", a)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World! Let's Gooo!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
