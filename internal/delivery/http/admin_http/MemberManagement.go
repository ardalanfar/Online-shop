package admin_http

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/dto/admin_dto"
	"Farashop/internal/service/admin_service"
	"Farashop/pkg/customerror"
	"encoding/json"

	"net/http"

	"github.com/labstack/echo/v4"
)

func ShowMembers(conn store.DbConn) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req = admin_dto.ShowMembersRequest{}

		//send service
		resp, err := admin_service.New(conn).ShowMembers(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, customerror.Unsuccessful())
		}

		//return ui
		return c.JSON(http.StatusOK, resp)
	}
}

func DeleteMember(conn store.DbConn) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req = admin_dto.DeleteMemberRequest{}

		//bind user information
		if errbind := json.NewDecoder(c.Request().Body).Decode(&req); errbind != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errbind.Error())
		}

		//validat information

		//send service

		//return ui
		return c.NoContent(http.StatusOK)
	}
}
