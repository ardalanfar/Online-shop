package admin

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/dto/admin"
	"Farashop/internal/service/adminservice"
	"Farashop/pkg/customerror"

	"net/http"

	"github.com/labstack/echo/v4"
)

func ShowMembers(conn store.DbConn) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req = admin.ShowMembersRequest{}

		resp, err := adminservice.New(conn).ShowMembers(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, customerror.Unsuccessful())
		}

		return c.JSON(http.StatusOK, resp)
	}
}

// func DeleteMember() {

// }
