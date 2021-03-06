package member_http

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/dto"
	"Farashop/internal/service/member_service"
	"Farashop/pkg/auth"
	"Farashop/pkg/customerror"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ShowOrders(conn store.DbConn) echo.HandlerFunc {
	return func(c echo.Context) error {

		u := c.Get("user").(*jwt.Token)
		claims := u.Claims.(*auth.Claims)

		var req = dto.ShowOrdersRequest{
			ID: claims.ID,
		}

		//Service
		resService, errService := member_service.NewMember(conn).ShowOrders(c.Request().Context(), req)
		if errService != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, customerror.Unsuccessful())
		}
		//return ui
		return c.JSON(http.StatusOK, resService)
	}
}
