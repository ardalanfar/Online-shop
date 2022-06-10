package auth

import (
	"Farashop/internal/entity"
	"Farashop/pkg/customerror"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

const (
	accessTokenCookieName  = "access-token"
	jwtSecretKey           = "some-secret-key"
	refreshTokenCookieName = "refresh-token"
	jwtRefreshSecretKey    = "some-refresh-secret-key"
)

type Claims struct {
	Name   string `json:"name"`
	ID     uint   `json:"id"`
	Access uint   `json:"Access"`
	jwt.StandardClaims
}

func GetJWTSecret() string {
	return jwtSecretKey
}

func GetRefreshJWTSecret() string {
	return jwtRefreshSecretKey
}

func GetAccessTokenCookieName() string {
	return accessTokenCookieName
}

func GetrefReshTokenCookieName() string {
	return refreshTokenCookieName
}

func GenerateTokensAndSetCookies(user entity.User, c echo.Context) error {
	accessToken, exp, err := generateAccessToken(user)
	if err != nil {
		return err
	}
	setTokenCookie(GetAccessTokenCookieName(), accessToken, exp, c)

	//generate here a new refresh token and saving it to the cookie.
	refreshToken, exp, err := generateRefreshToken(user)
	if err != nil {
		return err
	}
	setTokenCookie(GetrefReshTokenCookieName(), refreshToken, exp, c)
	return nil
}

func generateAccessToken(user entity.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

func generateToken(user entity.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	claims := &Claims{
		Name:   user.Username,
		ID:     user.ID,
		Access: user.Access,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Create the JWT string.
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
	//Http-only
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}

func JWTErrorChecker(err error, c echo.Context) error {
	return c.JSON(http.StatusForbidden, customerror.NOAccess())
}

func generateRefreshToken(user entity.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	return generateToken(user, expirationTime, []byte(GetRefreshJWTSecret()))
}
