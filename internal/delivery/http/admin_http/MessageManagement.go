package admin_http

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/service/admin_service"
	"Farashop/pkg/auth"
	"Farashop/pkg/customerror"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func SendMessage(conn store.DbConn, validator contract.ValidateSendMessage) echo.HandlerFunc {
	return func(c echo.Context) error {

		//bind user information
		userID, errbind := strconv.Atoi(c.Param("id"))
		if errbind != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errbind.Error())
		}
		var req = dto.SendMsgRequest{ID: uint(userID)}

		//validat information
		if errvalidat := validator(c.Request().Context(), req); errvalidat != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, customerror.InfoNotValid())
		}

		//send service
		resService, errservice := admin_service.NewMsg(conn).SendMsg(c.Request().Context(), req)
		if errservice != nil && resService.Result == false {
			return echo.NewHTTPError(http.StatusInternalServerError, customerror.InfoIncorrect())
		}

		//send service
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*auth.Claims)
		name := claims.Access

		//return ui
		return c.JSON(http.StatusOK, name)
	}
}
