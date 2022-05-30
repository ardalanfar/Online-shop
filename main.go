package main

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/delivery/http/admin_http"
	"Farashop/internal/delivery/http/public_http"
	"Farashop/internal/pkg/validator"

	"github.com/labstack/echo/v4"
)

func main() {

	//Connect to database and auto migrate
	conn := store.New()

	//Setup http server
	e := echo.New()

	//Routes
	e.POST("/register", public_http.Register(conn, validator.ValidateCreateUser(conn)))
	e.POST("/login", public_http.LoginUser(conn, validator.ValidateLoginUser(conn)))

	//Admin Group
	adminGroup := e.Group("/admin")
	/*----------------------MemberManagement----------------------*/
	MemberManagement := adminGroup.Group("/memberanagement")
	MemberManagement.GET("/showmembers", admin_http.ShowMembers(conn))
	MemberManagement.DELETE("/deletemember", admin_http.DeleteMember(conn, validator.ValidateDeleteMember(conn)))
	/*------------------------------------------------------------*/

	//Member Group
	//memberGroup := e.Group("/member")
	/*------------------------------------------------------------*/
	/*------------------------------------------------------------*/

	//Starting the server
	e.Logger.Fatal(e.Start(":8052"))
}
