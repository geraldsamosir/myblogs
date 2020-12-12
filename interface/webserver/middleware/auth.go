package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/geraldsamosir/myblogs/helper"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Auth struct {
	jwt.StandardClaims
	Data interface{} `json:"data"`
}

type HttpMethod string

const (
	POST   HttpMethod = "POST"
	GET    HttpMethod = "GET"
	PUT    HttpMethod = "PUT"
	PATCH  HttpMethod = "PATCH"
	DELETE HttpMethod = "DELETE"
	OPTION HttpMethod = "OPTION"
)

func (e HttpMethod) toString() string {
	switch e {
	case POST:
		return "POST"
	case GET:
		return "GET"
	case PUT:
		return "PUT"
	case PATCH:
		return "PATCH"
	case DELETE:
		return "DELETE"
	case OPTION:
		return "OPTION"
	default:
		return "METHOD NOT FOUND"
	}
}

type RouterAction struct {
	Url    string
	Method HttpMethod
}

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

func FindAllowRoutes(routes []RouterAction, url string, Method string) (int, bool) {
	for i, item := range routes {
		if strings.Contains(url, item.Url) && HttpMethod.toString(item.Method) == Method {
			return i, true
		}
	}
	return -1, false
}

func (auth *Auth) MiddlewareAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// allow routes
		routes := []RouterAction{
			RouterAction{
				"/api/Users/Login",
				"POST",
			},
			RouterAction{
				"/api/Users",
				"GET",
			},
			RouterAction{
				"/api/Articles",
				"GET",
			},
			RouterAction{
				"/api/Roles",
				"GET",
			},
			RouterAction{
				"/api/",
				"GET",
			},
			RouterAction{
				"/api",
				"GET",
			},
		}
		_, found := FindAllowRoutes(routes, c.Request().RequestURI, c.Request().Method)

		if found {
			return next(c)
		} else {
			authorizationHeader := c.Request().Header.Get("Authorization")
			var splitToken []string

			splitToken = strings.Split(authorizationHeader, "Bearer ")
			reqToken := splitToken[len(splitToken)-1]
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
