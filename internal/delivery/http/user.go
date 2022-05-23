package http

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/pkg/auth"
	"Farashop/internal/service/user"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(conn store.DbConn, validator contract.ValidateCreateUser) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req = dto.CreateUserRequest{}

		//bind user information
		if errbind := json.NewDecoder(c.Request().Body).Decode(&req); errbind != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errbind.Error())
		}

		//validat information
		if errvalidat := validator(req); errvalidat != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, errvalidat.Error())
		}

		//send service
		resp, errservice := user.New(conn).Register(c.Request().Context(), req)
		if errservice != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, errservice.Error())
		}

		//return ui
		return c.JSON(http.StatusOK, resp)
	}
}

func LoginUser(conn store.DbConn, validator contract.ValidateLoginUser) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req = dto.LoginUserRequest{}

		//bind user information
		if errbind := json.NewDecoder(c.Request().Body).Decode(&req); errbind != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errbind.Error())
		}

		//validat information
		if errvalidat := validator(req); errvalidat != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, errvalidat.Error())
		}

		//send service
		resp, errservice := user.New(conn).Login(c.Request().Context(), req)
		if errservice != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, errservice.Error())
		}

		//call pkg auth(create token and cookie)
		if resp.Result {
			err := auth.GenerateTokensAndSetCookies(resp, c)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Wellcom"+resp.User.Username)
			}
		}

		//return ui
		return c.JSON(http.StatusOK, resp)
	}
}
