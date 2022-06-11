package middlewares

import (
	"Farashop/internal/entity"
	"Farashop/pkg/auth"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetAdminGroup(grp *echo.Group) {

	grp.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &auth.Claims{},
		SigningKey:              []byte(auth.GetJWTSecret()),
		TokenLookup:             "cookie:access-token", // "<source>:<name>"
		ErrorHandlerWithContext: auth.JWTErrorChecker,
	}))

	grp.Use(TokenRefresherMiddleware)
}

func TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("user") == nil {
			return next(c)
		}
		u := c.Get("user").(*jwt.Token)
		claims := u.Claims.(*auth.Claims)

		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 10*time.Minute {
			rc, err := c.Cookie(auth.GetrefReshTokenCookieName())

			if err == nil && rc != nil {

				tkn, err := jwt.ParseWithClaims(rc.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(auth.GetRefreshJWTSecret()), nil
				})

				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						c.Response().Writer.WriteHeader(http.StatusUnauthorized)
					}
				}

				if tkn != nil && tkn.Valid {
					_ = auth.GenerateTokensAndSetCookies(entity.User{
						Username: claims.Name,
					}, c)
				}

			}
		}
		return next(c)
	}
}
