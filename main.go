package main

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/delivery/http"
	"Farashop/internal/pkg/validator"

	"github.com/labstack/echo/v4"
)

func main() {

	//connect to database and auto migrate
	conn := store.New()

	//setup http server
	e := echo.New()

	//add routes
	e.POST("/register", http.CreateUser(conn, validator.ValidateCreateUser))
	e.POST("/login", http.LoginUser(conn, validator.ValidateLoginUser))

	//run
	e.Logger.Fatal(e.Start(":8091"))
}
