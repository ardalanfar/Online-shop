package http

import (
	"Farashop/internal/adapter/store"
	"Farashop/internal/dto"
	"Farashop/internal/service/account"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(conn store.DbConn) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req = dto.CreateUserRequest{}

		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// if err := validator(req); err != nil {
		// 	return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		// }

		resp, err := account.New(conn).Register(c.Request().Context(), req)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}
