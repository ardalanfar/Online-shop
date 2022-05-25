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

// GenerateTokensAndSetCookies generates jwt token and saves it to the http-only cookie.
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
	cookie.Name = "user"
	cookie.Value = user.User.Username
	cookie.Expires = expiration
	cookie.Path = "/"
	c.SetCookie(cookie)
}

// func JWTErrorChecker() {
// }
