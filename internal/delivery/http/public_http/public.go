package public_http

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/config"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/service/public_service"
	"Farashop/pkg/auth"
	"Farashop/pkg/customerror"
	"Farashop/pkg/sendmsg"
	"encoding/json"
	"net/http"
	"strconv"

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
		if errBind := json.NewDecoder(c.Request().Body).Decode(&req); errBind != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errBind.Error())
		}
		//validat information
		if errValidat := validator(c.Request().Context(), req); errValidat != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, customerror.InfoNotValid())
		}
		//send service
		resService, errService := public_service.NewAuth(conn).Register(c.Request().Context(), req)
		if errService != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, customerror.InfoIncorrect())
		}
		//send email welcome register
		if resService.Result == true {
			config := config.GetConfig().Email
			msg := "Welcom" + resService.User.Username + "verify code :" + strconv.FormatUint(uint64(resService.User.Verification_code), 10)
			to := []string{
				resService.User.Email,
			}
			err := sendmsg.SendEmailUser(msg, to, config)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, customerror.SendEmailErr())
			}
		}
		//return ui
		return c.JSON(http.StatusOK, customerror.Successfully())
	}
}

func LoginUser(conn store.DbConn, validator contract.ValidateLoginUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req = dto.LoginUserRequest{}

		//bind user information
		if errBind := json.NewDecoder(c.Request().Body).Decode(&req); errBind != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errBind.Error())
		}
		//validat information
		if errValidat := validator(c.Request().Context(), req); errValidat != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, customerror.InfoNotValid())
		}
		//send service
		resService, errService := public_service.NewAuth(conn).Login(c.Request().Context(), req)
		if errService != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, customerror.InfoIncorrect())
		}
		//call pkg auth(create token and cookie)
		if resService.Result {
			err := auth.GenerateTokensAndSetCookies(resService.User, c)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, customerror.Unsuccessful())
			}
		}
		//return ui
		return echo.NewHTTPError(http.StatusUnauthorized, "Wellcom "+resService.User.Username)
	}
}
