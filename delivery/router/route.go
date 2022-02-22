package router

import (
	"net/http"
	"sirclo/delivery/controllers/auth"
	"sirclo/delivery/controllers/checkinandout"
	"sirclo/delivery/controllers/office"
	"sirclo/delivery/controllers/schedule"
	"sirclo/delivery/controllers/user"
	"sirclo/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(
	e *echo.Echo,
	userController *user.UserController,
	authController *auth.AuthController,
	scheduleController *schedule.ScheduleController,
	officeController *office.OfficeController,
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
	// schedule
	e.POST("/schedules", scheduleController.CreateSchedule(secret), middlewares.JWTMiddleware())
	e.PUT("/schedules/:id", scheduleController.EditSchedule(secret), middlewares.JWTMiddleware())
	e.GET("/schedules/:id", scheduleController.GetSchedule(secret), middlewares.JWTMiddleware())
	// office
	e.GET("/offices", officeController.GetOffices(secret), middlewares.JWTMiddleware())
	e.GET("/offices/:id", officeController.GetOffice(secret), middlewares.JWTMiddleware())
	//check in and out
	e.PUT("/checkin", checkController.Checkin(secret), middlewares.JWTMiddleware())
	e.PUT("/checkout", checkController.Checkout(secret), middlewares.JWTMiddleware())
}
