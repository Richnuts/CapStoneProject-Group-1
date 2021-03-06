package schedule

import (
	"math"
	"net/http"
	"sirclo/delivery/common"
	"sirclo/delivery/middlewares"
	"sirclo/entities"
	scheduleRepo "sirclo/repository/schedule"
	"strconv"
	"time"

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
		// check schedule
		kosong, _ := sr.repository.GetSchedulesByMonthAndYear(int(scheduleRequest.Month), scheduleRequest.Year, scheduleRequest.OfficeId)
		if kosong != nil {
			return c.JSON(http.StatusInternalServerError, common.CustomResponse(500, "failed creating schedule", "duplicate entry"))
		}
		// create schedule
		err_schedule := sr.repository.CreateSchedule(scheduleRequest.Month, scheduleRequest.Year, scheduleRequest.TotalCapacity, scheduleRequest.OfficeId)
		if err_schedule != nil {
			return c.JSON(http.StatusInternalServerError, common.CustomResponse(500, "failed creating schedule", "server error"))
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
		// getting the id
		scheduleId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		var scheduleEdit ScheduleEditFormat
		// prosess binding text
		if err_bind := c.Bind(&scheduleEdit); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		// mengedit schedule
		err_edit := sr.repository.EditSchedule(scheduleId, scheduleEdit.TotalCapacity)
		if err_edit != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation("berhasil mengedit jadwal"))
	}
}

func (sr ScheduleController) GetSchedule(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// getting the id
		scheduleId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		// getting the page
		pageString := c.QueryParam("page")
		halaman, err := strconv.Atoi(pageString)
		if err != nil {
			halaman = 1
		}
		offset := (halaman - 1) * 10
		// mengGet schedule
		var data entities.ScheduleResponse
		data, err_get := sr.repository.GetSchedule(scheduleId, offset)
		if err_get != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		gmt, _ := time.LoadLocation("Asia/Jakarta")
		data.Date = data.Date.In(gmt)
		// menGet total page
		data.TotalData, _ = sr.repository.GetTotalData(scheduleId)
		data.TotalPage = int((math.Ceil(float64(data.TotalData) / float64(10))))
		return c.JSON(http.StatusOK, data)
	}
}

func (sr ScheduleController) GetSchedulesByMonthAndYear(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// getting the month
		monthString := c.QueryParam("month")
		month, err := strconv.Atoi(monthString)
		if err != nil {
			return c.JSON(http.StatusForbidden, common.CustomResponse(400, "masukin woi bulannya", "bulan gaboleh kosong"))
		}
		// getting the year
		yearString := c.QueryParam("year")
		year, err := strconv.Atoi(yearString)
		if err != nil {
			return c.JSON(http.StatusForbidden, common.CustomResponse(400, "masukin woi tahunnya", "tahun gaboleh kosong"))
		}
		// getting the year
		officeId, err := strconv.Atoi(c.QueryParam("office"))
		if err != nil {
			return c.JSON(http.StatusForbidden, common.CustomResponse(400, "masukin woi officenya", "office gaboleh kosong"))
		}
		// mengGet schedule
		data, err_get := sr.repository.GetSchedulesByMonthAndYear(month, year, officeId)
		if err_get != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, data)
	}
}
