package main

import (
	"github.com/labstack/echo/v4"
)

func Less(a, b interface{}) bool {
	return a.(string) > b.(string)
}
func T_main() {

	e := echo.New()
	e.Static("/", ".")
	//	e.Logger.Fatal(e.Start(":9988"))
}
