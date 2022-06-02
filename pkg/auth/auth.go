package auth

import (
	"Farashop/internal/dto"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

//Intractor package auth user

const (
	accessTokenCookieName = "access-token"
	jwtSecretKey          = "some-secret-key"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func GetJWTSecret() string {
	return jwtSecretKey
}

// func GetRefreshJWTSecret() string {
// 	return jwtRefreshSecretKey
// }

type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func GenerateTokensAndSetCookies(user dto.LoginUserResponse, c echo.Context) error {
	accessToken, exp, err := generateAccessToken(user)
	if err != nil {
		return err
	}
	setTokenCookie(accessTokenCookieName, accessToken, exp, c)
	setUserCookie(user, exp, c)
	return nil
}

func generateAccessToken(user dto.LoginUserResponse) (string, time.Time, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

func generateToken(user dto.LoginUserResponse, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	claims := &jwtCustomClaims{
		Name:  user.User.Username,
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string.
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}
	return tokenString, expirationTime, nil
}

func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	// Http-only helps mitigate the risk of client side script accessing the protected cookie.
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}

func setUserCookie(user dto.LoginUserResponse, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = user.User.Username
	cookie.Expires = expiration
	cookie.Path = "/"
	c.SetCookie(cookie)
}

func JWTErrorChecker(err error, c echo.Context) error {
	// Redirects to the signIn form.
	return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("/register"))
}

// TokenRefresherMiddleware middleware, which refreshes JWT tokens if the access token is about to expire.
// func TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// If the user is not authenticated (no user token data in the context), don't do anything.
// 		if c.Get("username") == nil {
// 			return next(c)
// 		}
// 		// Gets user token from the context.
// 		u := c.Get("username").(*jwt.Token)

// 		claims := u.Claims.(*Claims)

// 		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 15*time.Minute {
// 			// Gets the refresh token from the cookie.
// 			rc, err := c.Cookie(refreshTokenCookieName)
// 			if err == nil && rc != nil {
// 				// Parses token and checks if it valid.
// 				tkn, err := jwt.ParseWithClaims(rc.Value, claims, func(token *jwt.Token) (interface{}, error) {
// 					return []byte(GetRefreshJWTSecret()), nil
// 				})
// 				if err != nil {
// 					if err == jwt.ErrSignatureInvalid {
// 						c.Response().Writer.WriteHeader(http.StatusUnauthorized)
// 					}
// 				}

// 				if tkn != nil && tkn.Valid {
// 					// If everything is good, update tokens.
// 					_ = GenerateTokensAndSetCookies(&username.User{
// 						Name: claims.Name,
// 					}, c)
// 				}
// 			}
// 		}
// 		return next(c)
// 	}
// }
