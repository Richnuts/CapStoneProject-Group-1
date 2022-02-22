package main

import (
	"log"
	"os"
	"sirclo/config"
	_authController "sirclo/delivery/controllers/auth"
	_scheduleController "sirclo/delivery/controllers/schedule"
	_userController "sirclo/delivery/controllers/user"
	"sirclo/delivery/router"
	_authRepo "sirclo/repository/auth"
	_scheduleRepo "sirclo/repository/schedule"
	_userRepo "sirclo/repository/user"
	"sirclo/util"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("SECRET")
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	db := util.MysqlDriver(config)

	//initiate user model
	userRepo := _userRepo.New(db)
	authRepo := _authRepo.New(db)
	scheduleRepo := _scheduleRepo.New(db)

	//initiate user controller
	userController := _userController.New(userRepo)
	authController := _authController.New(authRepo)
	scheduleController := _scheduleController.New(scheduleRepo)
	//create echo http
	e := echo.New()

	//register API path and controller
	router.RegisterPath(e,
		userController,
		authController,
		scheduleController,
		secret,
	)
	// run server
	e.Logger.Fatal(e.Start(":" + os.Getenv("Port")))
}
