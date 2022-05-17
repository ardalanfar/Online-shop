package main

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/delivery/http"

	"github.com/labstack/echo/v4"
)

func main() {

	//connect to database and auto migrate
	conn := store.New()

	//setup http server
	e := echo.New()

	//add routes
	e.POST("/register", http.CreateUser(conn))
	e.POST("/login", http.LoginUser(conn))

	//run
	e.Logger.Fatal(e.Start(":8010"))
}
