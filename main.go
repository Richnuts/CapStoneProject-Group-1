package main

import (
	"fmt"
	"log"
	"os"
	"sirclo/config"
	_attendanceController "sirclo/delivery/controllers/attendance"
	_authController "sirclo/delivery/controllers/auth"
	_certificateController "sirclo/delivery/controllers/certificate"
	_checkController "sirclo/delivery/controllers/checkinandout"
	_officeController "sirclo/delivery/controllers/office"
	_scheduleController "sirclo/delivery/controllers/schedule"
	_userController "sirclo/delivery/controllers/user"
	"time"

	"sirclo/delivery/router"
	_attendanceRepo "sirclo/repository/attendance"
	_authRepo "sirclo/repository/auth"
	_certificateRepo "sirclo/repository/certificate"
	_checkRepo "sirclo/repository/checkinandout"
	_officeRepo "sirclo/repository/office"
	_scheduleRepo "sirclo/repository/schedule"
	_userRepo "sirclo/repository/user"
	"sirclo/util"

	"github.com/go-co-op/gocron"
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
	// cron job
	gmt, _ := time.LoadLocation("Asia/Jakarta")
	s := gocron.NewScheduler(gmt)
	s.Every(1).Day().At("00:00").Do(
		func() {
			_, err := db.Exec(`
			UPDATE 
				attendances 
			SET
				check_out = now() 
			WHERE 
				check_in is not null AND DAY(CONVERT_TZ(check_in, '+00:00', '+7:00')) = ? AND check_out is null`, (time.Now().Day())-1)
			if err != nil {
				fmt.Println("gagal auto logout")
			} else {
				fmt.Println("autologout berjalan")
			}
		},
	)
	s.StartAsync()
	//initiate user model
	userRepo := _userRepo.New(db)
	authRepo := _authRepo.New(db)
	scheduleRepo := _scheduleRepo.New(db)
	officeRepo := _officeRepo.New(db)
	checkRepo := _checkRepo.New(db)
	certificateRepo := _certificateRepo.New(db)
	attendanceRepo := _attendanceRepo.New(db)

	//initiate user controller
	userController := _userController.New(userRepo)
	authController := _authController.New(authRepo)
	scheduleController := _scheduleController.New(scheduleRepo)
	officeController := _officeController.New(officeRepo)
	checkController := _checkController.New(checkRepo)
	certificateController := _certificateController.New(certificateRepo)
	attendanceController := _attendanceController.New(attendanceRepo)

	//create echo http
	e := echo.New()

	//register API path and controller
	router.RegisterPath(e,
		userController,
		authController,
		scheduleController,
		officeController,
		checkController,
		certificateController,
		attendanceController,
		secret,
	)
	// run server
	e.Logger.Fatal(e.Start(":" + os.Getenv("Port")))
}
