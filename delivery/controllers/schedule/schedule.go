package schedule

import (
	"fmt"
	"net/http"
	"sirclo/delivery/common"
	"sirclo/delivery/middlewares"
	scheduleRepo "sirclo/repository/schedule"

	"github.com/labstack/echo/v4"
)

type ScheduleController struct {
	repository scheduleRepo.Schedule
}

func New(schedule scheduleRepo.Schedule) *ScheduleController {
	return &ScheduleController{
		repository: schedule,
	}
}

func (sr ScheduleController) CreateSchedule(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// check role
		role := middlewares.GetUserRole(secret, c)
		if role != "admin" {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		var scheduleRequest ScheduleRequestFormat
		// prosess binding text
		if err_bind := c.Bind(&scheduleRequest); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		if scheduleRequest.TotalCapacity == 0 {
			scheduleRequest.TotalCapacity = 50
		}
		fmt.Println(scheduleRequest)
		// create schedule
		err_schedule := sr.repository.CreateSchedule(scheduleRequest.Month, scheduleRequest.Year, scheduleRequest.TotalCapacity, scheduleRequest.OfficeId)
		if err_schedule != nil {
			fmt.Println(err_schedule)
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "failed creating schedule", "duplicate entry"))
		}
		return c.JSON(http.StatusOK, common.SuccessOperation("berhasil membuat jadwal"))
	}
}

func (sr ScheduleController) EditSchedule(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// check role
		role := middlewares.GetUserRole(secret, c)
		if role != "admin" {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		var scheduleEdit ScheduleEditFormat
		// prosess binding text
		if err_bind := c.Bind(&scheduleEdit); err_bind != nil {
			fmt.Println(err_bind)
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		// mengedit schedule
		err_edit := sr.repository.EditSchedule(scheduleEdit.Date, scheduleEdit.TotalCapacity, scheduleEdit.OfficeId)
		if err_edit != nil {
			fmt.Println(err_edit)
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation("berhasil mengedit jadwal"))
	}
}

// gmt := time.FixedZone("gmt+7", +7*60*60)
// end := time.Now().In(gmt)
// fmt.Println("ini end = ", end)
// start := time.Date(2022, 2, 27, 0, 0, 0, 0, gmt)
// diff := start.Sub(end)
// fmt.Println("ini start = ", start)
// fmt.Println("ini end = ", end)
// fmt.Println("ini diff = ", diff)
// //time Since
// fmt.Println(time.Since(start.Add(-time.Hour * 1)))
// waktucantik := end.Format(time.RFC850)
// fmt.Println(waktucantik)
