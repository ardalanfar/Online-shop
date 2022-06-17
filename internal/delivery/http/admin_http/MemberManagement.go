package admin_http

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/service/admin_service"
	"Farashop/pkg/customerror"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ShowMembers(conn store.DbConn) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req = dto.ShowMembersRequest{}

		//send service
		resService, errService := admin_service.NewAdmin(conn).ShowMembers(c.Request().Context(), req)
		if errService != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, customerror.Unsuccessful())
		}
		//return ui
		return c.JSON(http.StatusOK, resService)
	}
}

func DeleteMember(conn store.DbConn, validator contract.ValidateDeleteMember) echo.HandlerFunc {
	return func(c echo.Context) error {

		//bind user information
		userID, errBind := strconv.Atoi(c.Param("id"))
		if errBind != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errBind.Error())
		}
		var req = dto.DeleteMemberRequest{ID: uint(userID)}
		//validat information
		if errValidat := validator(c.Request().Context(), req); errValidat != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, customerror.InfoNotValid())
		}
		//send service
		resService, errService := admin_service.NewAdmin(conn).DeleteMember(c.Request().Context(), req)
		if errService != nil && resService.Result == false {
			return echo.NewHTTPError(http.StatusInternalServerError, customerror.InfoIncorrect())
		}
		//return ui
		return c.JSON(http.StatusOK, customerror.Successfully())
	}
}

func ShowInfoMember(conn store.DbConn, validator contract.ValidateShowInfoMember) echo.HandlerFunc {
	return func(c echo.Context) error {

		//bind user information
		userID, errBind := strconv.Atoi(c.Param("id"))
		if errBind != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errBind.Error())
		}
		var req = dto.ShowInfoMemberRequest{ID: uint(userID)}

		//validat information
		if errValidat := validator(c.Request().Context(), req); errValidat != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, customerror.InfoNotValid())
		}
		//send service
		resService, errService := admin_service.NewAdmin(conn).ShowInfoMember(c.Request().Context(), req)
		if errService != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, resService)
		}
		//return ui
		return c.JSON(http.StatusOK, resService)
	}
}
