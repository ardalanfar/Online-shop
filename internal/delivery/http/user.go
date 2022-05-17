package http

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/service/user"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(conn store.DbConn, validator contract.ValidateCreateUser) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req = dto.CreateUserRequest{}

		if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// //validat
		if err := validator(req); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		//send usecase
		resp, err := user.New(conn).Register(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		//return ui
		return c.JSON(http.StatusOK, resp)
	}
}

func LoginUser(conn store.DbConn) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req = dto.LoginUserRequest{}

		if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		//validat

		//send service
		resp, err := user.New(conn).Login(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		//return ui
		return c.JSON(http.StatusOK, resp)
	}
}
