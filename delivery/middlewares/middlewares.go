package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("rahasia"),
	})
}

func GetUserId(secret string, e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user != nil && user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))
		return userId
	}
	return 0
}

func CreateToken(id int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 144).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("rahasia"))
}
