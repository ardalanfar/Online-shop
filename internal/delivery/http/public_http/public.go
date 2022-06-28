package public_http

import (
	"Farashop/internal/adapter/sendmsg"
	"Farashop/internal/adapter/store"
	"Farashop/internal/config"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/service/public_service"
	"Farashop/pkg/auth"
	"Farashop/pkg/customerror"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(conn store.DbConn, validator contract.ValidateRegister) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req = dto.RegisterUserRequest{}

		//bind user information
		if errBind := json.NewDecoder(c.Request().Body).Decode(&req); errBind != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errBind.Error())
		}
		//validat information
		if errValidat := validator(c.Request().Context(), req); errValidat != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, customerror.InfoNotValid())
		}
		//service
		resService, errService := public_service.NewPublic(conn).Register(c.Request().Context(), req)
		if errService != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, customerror.InfoIncorrect())
		}
		//send email welcome register
		if resService.Result {
			config := config.GetConfig().Email
			to := []string{
				req.Email,
			}
			subject := "verify register"
			request := sendmsg.Mail{
				Sender:  config.From,
				To:      to,
				Subject: subject,
			}
			//build message
			msg := request.BuildMessage()
			//send email
			err := request.SendEmailUser(msg, config)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, customerror.SendEmailErr())
			}
		}
		//return ui
		return c.JSON(http.StatusOK, customerror.Successfully())
	}
}

func Login(conn store.DbConn, validator contract.ValidateLogin) echo.HandlerFunc {
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
		//service
		resService, errService := public_service.NewPublic(conn).Login(c.Request().Context(), req)
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
		return echo.NewHTTPError(http.StatusOK, "Wellcom "+resService.User.Username)
	}
}

func MemberValidation(conn store.DbConn, validator contract.ValidateMemberValidation) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req = dto.MemberValidationRequest{}

		//bind user information
		if errBind := json.NewDecoder(c.Request().Body).Decode(&req); errBind != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errBind.Error())
		}
		//validat information
		if errValidat := validator(c.Request().Context(), req); errValidat != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, customerror.InfoNotValid())
		}
		//service
		resService, errService := public_service.NewPublic(conn).MemberValidation(c.Request().Context(), req)
		if errService != nil || !resService.Result {
			return echo.NewHTTPError(http.StatusInternalServerError, customerror.InfoIncorrect())
		}
		//return ui
		return c.JSON(http.StatusOK, customerror.Successfully())

	}
}
