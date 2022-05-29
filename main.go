package main

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/delivery/http"
	"Farashop/internal/delivery/http/admin"
	"Farashop/internal/pkg/validator"

	"github.com/labstack/echo/v4"
)

func main() {

	//connect to database and auto migrate
	conn := store.New()

	//setup http server
	e := echo.New()

	//add routes
	e.POST("/register", http.CreateUser(conn, validator.ValidateCreateUser(conn)))
	e.POST("/login", http.LoginUser(conn, validator.ValidateLoginUser(conn)))

	//Admin Group
	adminGroup := e.Group("/admin")

	/*------------------------------------------------------------*/
	MemberManagement := adminGroup.Group("/MemberManagement")
	MemberManagement.GET("/ShowMembers", admin.ShowMembers(conn))

	/*------------------------------------------------------------*/

	/*------------------------------------------------------------*/

	//userGroup := e.Group("/user")

	//starting the server
	e.Logger.Fatal(e.Start(":8031"))
}
