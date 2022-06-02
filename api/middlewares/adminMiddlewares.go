package middlewares

import (
	"Farashop/pkg/auth"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetAdminMiddleware(e *echo.Echo) {

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &auth.Claims{},
		SigningKey:              []byte(auth.GetJWTSecret()),
		TokenLookup:             "cookie:access-token", // "<source>:<name>"
		ErrorHandlerWithContext: auth.JWTErrorChecker,
	}))
}

func SetAdminGroup(grp *echo.Group) {
	grp.Use(checkCookie)
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(("username"))

		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "you dont have any cookie")
			}
			log.Panicln(err)
			return err
		}

		if cookie.Value == "admin" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "you dont have the right cookie")
	}
}
