package http

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/service/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(conn store.DbConn, validator contract.ValidateCreateUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req = dto.CreateUserRequest{}

		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		//validat
		if err := validator(req); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		//send service
		resp, err := user.New(conn).Register(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		//return
		return c.JSON(http.StatusOK, resp)
	}
}

func LoginUser(conn store.DbConn) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		password := c.Param("password")

		var req = dto.LoginUserRequest{}

		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		req.Password = password
		req.Username = username

		//validat
		// if err := validator(c.Request().Context(), req); err != nil {
		// 	return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		// }

		//send service
		resp, err := user.New(conn).Login(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		//return
		return c.JSON(http.StatusOK, resp)
	}
}
