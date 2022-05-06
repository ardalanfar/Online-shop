package main

import (
	"Farashop/internal/adapter/store"

	"github.com/labstack/echo/v4"
)

func main() {

	//connect to database and auto migrate
	conn := store.New()

	//setup http server and router
	e := echo.New()

	e.POST("/register",http.)
	_ = conn

	//add routes

}
