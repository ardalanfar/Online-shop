package main

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/delivery/http/admin_http"
	"Farashop/internal/delivery/http/public_http"
	"Farashop/internal/pkg/validator"

	"github.com/labstack/echo/v4"
)

func main() {

	//connect to database and auto migrate
	conn := store.New()

	//setup http server
	e := echo.New()

	//add routes
	e.POST("/register", public_http.CreateUser(conn, validator.ValidateCreateUser(conn)))
	e.POST("/login", public_http.LoginUser(conn, validator.ValidateLoginUser(conn)))

	//Admin Group
	adminGroup := e.Group("/admin")

	/*------------------------------------------------------------*/
	MemberManagement := adminGroup.Group("/MemberManagement")
	MemberManagement.GET("/showmembers", admin_http.ShowMembers(conn))
	//MemberManagement.DELETE("/deletmember", admin_dto.DeleteMember(conn,validator.ValidateDeleteMember)))

	/*------------------------------------------------------------*/

	/*------------------------------------------------------------*/

	//userGroup := e.Group("/user")

	//starting the server
	e.Logger.Fatal(e.Start(":8031"))
}
