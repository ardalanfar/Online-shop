package middlewares

import (
	"Farashop/internal/entity"
	"Farashop/pkg/auth"
	"Farashop/pkg/customerror"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetMemberGroup(grp *echo.Group) {

	grp.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &auth.Claims{},
		SigningKey:              []byte(auth.GetJWTSecret()),
		TokenLookup:             "cookie:" + auth.GetAccessTokenCookieName(), // "<source>:<name>"
		ErrorHandlerWithContext: auth.JWTErrorChecker,
	}))

	grp.Use(TokenRefresherMiddlewareMember)
	grp.Use(CheckAccessMember)
}

func CheckAccessMember(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("user") == nil {
			return next(c)
		}
		u := c.Get("user").(*jwt.Token)
		claims := u.Claims.(*auth.Claims)
		if int(claims.Access) == 2 {
			return next(c)
		}
		return c.JSON(http.StatusBadRequest, customerror.NOAccess())
	}
}

func TokenRefresherMiddlewareMember(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("user") == nil {
			return next(c)
		}
		u := c.Get("user").(*jwt.Token)
		claims := u.Claims.(*auth.Claims)

		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < (10 * time.Minute) {
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
						ID:       claims.ID,
						Access:   claims.Access,
					}, c)
				}
			}
		}
		return next(c)
	}
}
