package main

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/delivery/http"
	"Farashop/internal/validator"

	"github.com/labstack/echo/v4"
)

func main() {

	//connect to database and auto migrate
	connection := store.New()

	//setup http server
	e := echo.New()

	//add routes
	e.POST("/register", http.CreateUser(connection, validator.ValidateCreateUser))
	e.POST("/login", http.LoginUser(connection))

	//run
	e.Logger.Fatal(e.Start(":8080"))

}
