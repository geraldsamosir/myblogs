package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/geraldsamosir/myblogs/helper"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

type Auth struct {
	jwt.StandardClaims
	Data interface{} `json:"data"`
}

// Errors
var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusBadRequest, "missing or malformed jwt")
)

type (
	jwtExtractor func(echo.Context) (string, error)
)

func (*Auth) GenerateToken(data interface{}) (string, error) {
	claim := Auth{
		StandardClaims: jwt.StandardClaims{
			Issuer:    viper.GetString("JWT_ISSUER"),
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		Data: data,
	}
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := sign.SignedString([]byte(viper.GetString("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}

func FindAllowRoutes(routes []string, val string) (int, bool) {
	for i, item := range routes {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func (auth *Auth) MiddlewareAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// allow routes
		routes := []string{
			"/api/Users/Login",
		}

		_, found := FindAllowRoutes(routes, c.Request().RequestURI)

		if c.Request().Method == "GET" || found {
			return next(c)
		} else {
			authorizationHeader := c.Request().Header.Get("Authorization")
			var splitToken []string

			splitToken = strings.Split(authorizationHeader, "Bearer ")
			reqToken := splitToken[len(splitToken)-1]
			log.Println("sini", reqToken)

			token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
				return []byte(viper.GetString("JWT_SECRET")), nil
			})

			if err != nil {
				return helper.Response(http.StatusUnauthorized, nil, "token invalid", c)
			}

			_, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				return helper.Response(http.StatusUnauthorized, nil, "token invalid", c)

			}
			return next(c)
		}
	}

}

func InitMiddleware() *Auth {
	return &Auth{}
}
