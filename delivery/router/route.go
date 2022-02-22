package router

import (
	"net/http"
	"sirclo/delivery/controllers/auth"
	"sirclo/delivery/controllers/checkinandout"
	"sirclo/delivery/controllers/user"
	"sirclo/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(
	e *echo.Echo,
	userController *user.UserController,
	authController *auth.AuthController,
	checkController *checkinandout.CheckController,
	secret string,
) {
	// logger
	e.Pre(middleware.RemoveTrailingSlash(), middleware.Logger())
	// cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	// auth
	e.POST("/login", authController.Login())
	// e.POST("/register", authController.Register())
	// user
	e.GET("/profile", userController.GetProfile(secret), middlewares.JWTMiddleware())
	e.GET("/users/:id", userController.GetUser(secret), middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userController.DeleteUser(secret), middlewares.JWTMiddleware())
	e.PUT("/users/:id", userController.EditUser(secret), middlewares.JWTMiddleware())
	//check in and out
	e.PUT("/checkin", checkController.Checkin(secret), middlewares.JWTMiddleware())
	e.PUT("/checkout", checkController.Checkout(secret), middlewares.JWTMiddleware())
}
