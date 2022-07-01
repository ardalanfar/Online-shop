package main

import (
	"Farashop/api/middlewares"
	"Farashop/internal/adapter/store"
	"Farashop/internal/config"
	"Farashop/internal/delivery/http/admin_http"
	"Farashop/internal/delivery/http/member_http"
	"Farashop/internal/delivery/http/public_http"
	"Farashop/internal/pkg/validator"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	//Connect to database and auto migrate
	conn := store.New()

	//Setup http server
	e := echo.New()

	//Logger
	filePath, _ := os.OpenFile(config.GetConfig().Log.LogDirectory, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	e.Logger.SetOutput(filePath)
	defer filePath.Close()

	//Middlewares
	middlewares.SetMainMiddleware(e)

	//Routes
	/*--------------------------------------------------------------*/
	//Public
	e.POST("/register", public_http.Register(conn, validator.ValidateRegister(conn)))
	e.POST("/login", public_http.Login(conn, validator.ValidateLogin(conn)))
	e.PATCH("/validation", public_http.MemberValidation(conn, validator.ValidateMemberValidation(conn)))

	/*--------------------------------------------------------------*/
	//Admin Group
	AdminGroup := e.Group("/admin")
	middlewares.SetAdminGroup(AdminGroup)

	/*------------Member Management------------*/
	MemberManagement := AdminGroup.Group("/members")
	MemberManagement.GET("/showmembers", admin_http.ShowMembers(conn))
	MemberManagement.DELETE("/deletemember/:id", admin_http.DeleteMember(conn, validator.ValidateDeleteMember(conn)))
	MemberManagement.POST("/showinfo/:id", admin_http.ShowInfoMember(conn, validator.ValidateShowInfoMember(conn)))
	/*--------------------------------------------------------------*/

	//Member Group
	MemberGroup := e.Group("/member")
	middlewares.SetMemberGroup(MemberGroup)

	/*--------------Order Management-------------*/

	OrderManagement := MemberGroup.Group("/orders")
	OrderManagement.GET("/showorders", member_http.ShowOrders(conn))

	/*--------------------------------------------------------------*/

	//Starting the server
	e.Logger.Fatal(e.Start(":8014"))
}
