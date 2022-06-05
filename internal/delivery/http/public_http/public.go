package public_http

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/service/public_service"
	"Farashop/pkg/auth"
	"Farashop/pkg/customerror"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(conn store.DbConn, validator contract.ValidateCreateUser) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req = dto.CreateUserRequest{}

		//bind user information
		// if errbind := c.Bind(&req); errbind != nil {
		// 	return echo.NewHTTPError(http.StatusBadRequest, errbind.Error())
		// }

		//bind user information
		if errbind := json.NewDecoder(c.Request().Body).Decode(&req); errbind != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errbind.Error())
		}
		//validat information
		if errvalidat := validator(c.Request().Context(), req); errvalidat != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, customerror.InfoNotValid())
		}
		//send service
		_, errservice := public_service.NewAuth(conn).Register(c.Request().Context(), req)
		if errservice != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, customerror.InfoIncorrect())
		}
		//return ui
		return c.JSON(http.StatusOK, customerror.Successfully())
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
		if errvalidat := validator(c.Request().Context(), req); errvalidat != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, customerror.InfoNotValid())
		}
		//send service
		resp, errservice := public_service.NewAuth(conn).Login(c.Request().Context(), req)
		if errservice != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, customerror.InfoIncorrect())
		}
		//call pkg auth(create token and cookie)
		if resp.Result {
			err := auth.GenerateTokensAndSetCookies(resp.User, c)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, customerror.Unsuccessful())
			}
		}
		//return ui
		return echo.NewHTTPError(http.StatusUnauthorized, "Wellcom "+resp.User.Username)
	}
}
