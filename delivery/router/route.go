package router

import (
	"net/http"
	"sirclo/delivery/controllers/attendance"
	"sirclo/delivery/controllers/auth"
	"sirclo/delivery/controllers/certificate"
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
	certificateController *certificate.CertificateController,
	attendanceController *attendance.AttendanceController,
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
	e.GET("/schedules", scheduleController.GetSchedulesByMonthAndYear(secret), middlewares.JWTMiddleware())
	// office
	e.GET("/offices", officeController.GetOffices(secret), middlewares.JWTMiddleware())
	e.GET("/offices/:id", officeController.GetOffice(secret), middlewares.JWTMiddleware())
	//check in and out
	e.PUT("/checkin", checkController.Checkin(secret), middlewares.JWTMiddleware())
	e.PUT("/checkout", checkController.Checkout(secret), middlewares.JWTMiddleware())

	//certificate
	e.GET("/mycertificates", certificateController.GetMyCertificate(secret), middlewares.JWTMiddleware())
	e.GET("/certificates", certificateController.GetUsersCertificates(secret), middlewares.JWTMiddleware())
	e.GET("/certificates/:id", certificateController.GetCertificateById(secret), middlewares.JWTMiddleware())
	e.POST("/certificates", certificateController.CreateCertificate(secret), middlewares.JWTMiddleware())
	e.PUT("/certificates/:id", certificateController.EditCertificate(secret), middlewares.JWTMiddleware())
	e.PUT("/mycertificates/:id", certificateController.EditMyCertificate(secret), middlewares.JWTMiddleware())

	// attendance
	e.POST("/attendances", attendanceController.CreateAttendance(secret), middlewares.JWTMiddleware())
	e.PUT("/attendances/:id", attendanceController.EditAttendance(secret), middlewares.JWTMiddleware())
	e.GET("/attendances/:id", attendanceController.GetAttendanceById(secret), middlewares.JWTMiddleware())
	e.GET("/myattendances", attendanceController.GetMyAttendance(secret), middlewares.JWTMiddleware())
	e.GET("/mylatestattendances", attendanceController.GetMyAttendanceSortByLatest(secret), middlewares.JWTMiddleware())
	e.GET("/mylongestattendances", attendanceController.GetMyAttendanceSortByLongest(secret), middlewares.JWTMiddleware())
	e.GET("/pendingattendances", attendanceController.GetPendingAttendance(secret), middlewares.JWTMiddleware())
}
