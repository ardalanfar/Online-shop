package main

import (
	"Farashop/api/middlewares"
	"Farashop/internal/adapter/store"
	"Farashop/internal/delivery/http/admin_http"
	"Farashop/internal/delivery/http/public_http"
	"Farashop/internal/pkg/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	//Connect to database and auto migrate
	conn := store.New()

	//Setup http server
	e := echo.New()

	//Middlewares
	middlewares.SetMainMiddleware(e)
	e.Use(middleware.Recover())

	//Routes

	/*--------------------------------------------------------------*/
	//Public Group
	e.POST("/register", public_http.Register(conn, validator.ValidateRegister(conn)))
	e.POST("/login", public_http.Login(conn, validator.ValidateLogin(conn)))
	e.PATCH("/validation", public_http.MemberValidation(conn, validator.ValidateMemberValidation(conn)))

	/*--------------------------------------------------------------*/
	//Admin Group
	adminGroup := e.Group("/admin")
	middlewares.SetAdminGroup(adminGroup)

	/*------------Member Management------------*/
	MemberManagement := adminGroup.Group("/membermanagement")
	MemberManagement.GET("/showmembers", admin_http.ShowMembers(conn))
	MemberManagement.DELETE("/deletemember/:id", admin_http.DeleteMember(conn, validator.ValidateDeleteMember(conn)))
	//MemberManagement.POST("/showinfo/:id", admin_http.ShowInfoMember(conn, validator.ValidateShowInfoMember(conn)))
	/*--------------------------------------------------------------*/

	//Member Group
	//memberGroup := e.Group("/member")
	/*-----------------------Order Management-----------------------*/
	/*--------------------------------------------------------------*/

	//Starting the server
	e.Logger.Fatal(e.Start(":8086"))
}
