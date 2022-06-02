package admin_http

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/service/admin_service"
	"Farashop/pkg/customerror"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ShowMembers(conn store.DbConn) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req = dto.ShowMembersRequest{}

		//send service
		resService, errservice := admin_service.New(conn).ShowMembers(c.Request().Context(), req)
		if errservice != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, customerror.Unsuccessful())
		}
		//return ui
		return c.JSON(http.StatusOK, resService)
	}
}

func DeleteMember(conn store.DbConn, validator contract.ValidateDeleteMember) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req = dto.DeleteMemberRequest{}

		//bind user information
		if errbind := json.NewDecoder(c.Request().Body).Decode(&req); errbind != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errbind.Error())
		}
		//validat information
		if errvalidat := validator(c.Request().Context(), req); errvalidat != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, customerror.InfoNotValid())
		}
		//send service
		_, errservice := admin_service.New(conn).DeleteMember(c.Request().Context(), req)
		if errservice != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, customerror.InfoIncorrect())
		}
		//return ui
		return c.JSON(http.StatusOK, customerror.Successfully())
	}
}
