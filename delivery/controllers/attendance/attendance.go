package attendance

import (
	"math"
	"net/http"
	"sirclo/delivery/common"
	"sirclo/delivery/controllers/imageLib"
	"sirclo/delivery/middlewares"
	"sirclo/entities"
	attendanceRepo "sirclo/repository/attendance"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type AttendanceController struct {
	repository attendanceRepo.Attendance
}

func New(attendance attendanceRepo.Attendance) *AttendanceController {
	return &AttendanceController{
		repository: attendance,
	}
}

func (ac AttendanceController) CreateAttendance(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// check user status vaksin
		err_vaccine := ac.repository.GetUserVaccineStatus(loginId)
		if err_vaccine != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "operation failed", "user belum vaccine"))
		}
		var attendanceRequest AttendanceRequestFormat
		// prosess binding text
		if err_bind := c.Bind(&attendanceRequest); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "operation failed", "binding text gagal"))
		}
		// check tanggal request > tanggal sekarang
		requestDate, err_check_date := ac.repository.CheckCreateRequestDate(attendanceRequest.ScheduleId)
		if err_check_date != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "operation failed", "tanggal nya ga ada woi"))
		}
		gmt, _ := time.LoadLocation("Asia/Jakarta")
		currentDate := time.Now()
		checkDate := currentDate.Before(requestDate.In(gmt))
		if !checkDate {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "operation failed", "tanggal request harus lebih besar daripada tanggal hari ini"))
		}
		// Check requestnya udah ada belom
		err_checking := ac.repository.GetUserAttendanceStatus(loginId, attendanceRequest.ScheduleId)
		if err_checking != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "operation failed", "request telah ada"))
		}
		// proses binding image
		fileData, fileInfo, err_binding_image := c.Request().FormFile("image")
		if err_binding_image != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "operation failed", "binding image gagal"))
		}
		// check file extension
		_, err_check_extension := imageLib.CheckFileExtension(fileInfo.Filename)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "bad request", "file extension error"))
		}
		// check file size
		err_check_size := imageLib.CheckFileSize(fileInfo.Size)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "bad request", "file size error"))
		}
		fileName := "attendance_" + strconv.Itoa(attendanceRequest.ScheduleId) + "_" + strconv.Itoa(loginId)
		// upload the photo
		var err_upload_photo error
		theUrl, err_upload_photo := imageLib.UploadImage("attendance", fileName, fileData)
		if err_upload_photo != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "bad request", "Upload Image Failed"))
		}
		err_create := ac.repository.CreateAttendance(loginId, attendanceRequest.ScheduleId, attendanceRequest.Description, theUrl)
		if err_create != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(500, "internal server error", "Failed Creating Entity in database"))
		}
		return c.JSON(http.StatusOK, common.CustomResponse(200, "operation success", "berhasil membuat request WFO"))
	}
}

func (ac AttendanceController) EditAttendance(secret string) echo.HandlerFunc {
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
		var attendanceRequest AttendanceEditFormat
		// getting the id
		attendanceId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		// prosess binding text
		if err_bind := c.Bind(&attendanceRequest); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "operation failed", "binding text gagal"))
		}
		if attendanceRequest.StatusInfo == "" {
			attendanceRequest.StatusInfo = "-"
		}
		// check capacity
		capacity, err_capacity := ac.repository.CheckCapacity(attendanceRequest.ScheduleId)
		if err_capacity != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(500, "internal server error", "gagal mengecheck capacity"))
		}
		if capacity < 1 {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(500, "internal server error", "Kapasitas telah penuh"))
		}
		// edit the status
		err_edit := ac.repository.EditAttendance(attendanceId, loginId, attendanceRequest.Status, attendanceRequest.StatusInfo)
		if err_edit != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(500, "internal server error", "gagal merubah status / request sudah diedit"))
		}
		return c.JSON(http.StatusOK, common.CustomResponse(200, "operation success", "berhasil merubah status request WFO"))
	}
}

func (ac AttendanceController) GetAttendanceById(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// getting the id
		attendanceId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		// get the attendance
		hasil, err_get := ac.repository.GetAttendanceById(attendanceId)
		if err_get != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(500, "internal server error", "request tidak ditemukan"))
		}
		return c.JSON(http.StatusOK, hasil)
	}
}

func (ac AttendanceController) GetMyAttendance(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// getting the page
		pageString := c.QueryParam("page")
		halaman, err := strconv.Atoi(pageString)
		if err != nil {
			halaman = 1
		}
		offset := (halaman - 1) * 10
		// getting the status
		status := c.QueryParam("status")
		// get the attendance
		hasil, err_get := ac.repository.GetMyAttendance(loginId, offset, status)
		if err_get != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(500, "internal server error", "request tidak ditemukan"))
		}
		// get total page
		page, _ := ac.repository.GetMyAttendanceTotalPage(loginId, status)

		data := entities.AttendancePageFormat{TotalPage: page, Attendance: hasil}

		return c.JSON(http.StatusOK, data)
	}
}

func (ac AttendanceController) GetMyAttendanceSortByLatest(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// getting the page
		pageString := c.QueryParam("page")
		halaman, err := strconv.Atoi(pageString)
		if err != nil {
			halaman = 1
		}
		offset := (halaman - 1) * 10
		// getting the status
		status := c.QueryParam("status")
		// get the attendance
		hasil, err_get := ac.repository.GetMyAttendanceSortByLatest(loginId, offset, status)
		if err_get != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(500, "internal server error", "request tidak ditemukan"))
		}
		// get total page
		page, _ := ac.repository.GetMyAttendanceTotalPage(loginId, status)

		data := entities.AttendancePageFormat{TotalPage: page, Attendance: hasil}
		return c.JSON(http.StatusOK, data)
	}
}

func (ac AttendanceController) GetMyAttendanceSortByLongest(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// getting the page
		pageString := c.QueryParam("page")
		halaman, err := strconv.Atoi(pageString)
		if err != nil {
			halaman = 1
		}
		offset := (halaman - 1) * 10
		// getting the status
		status := c.QueryParam("status")
		// get the attendance
		hasil, err_get := ac.repository.GetMyAttendanceSortByLongest(loginId, offset, status)
		if err_get != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(500, "internal server error", "request tidak ditemukan"))
		}
		// get total page
		page, _ := ac.repository.GetMyAttendanceTotalPage(loginId, status)

		data := entities.AttendancePageFormat{TotalPage: page, Attendance: hasil}
		return c.JSON(http.StatusOK, data)
	}
}

func (ac AttendanceController) GetPendingAttendance(secret string) echo.HandlerFunc {
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
		// getting the page
		officeString := c.QueryParam("office")
		officeId, err_office := strconv.Atoi(officeString)
		if err_office != nil {
			return c.JSON(http.StatusBadGateway, common.CustomResponse(400, "Operation Failed", "Office harus diisi"))
		}
		// getting the page
		pageString := c.QueryParam("page")
		halaman, err := strconv.Atoi(pageString)
		if err != nil {
			halaman = 1
		}
		offset := (halaman - 1) * 10

		// get the attendance
		hasil, err_get := ac.repository.GetPendingAttendance(offset, officeId)
		if err_get != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(500, "internal server error", "request tidak ditemukan"))
		}
		// get total page
		countData, _ := ac.repository.GetPendingAttendanceTotalPage(officeId)

		data := entities.PendingAttendancePageFormat{TotalPage: int((math.Ceil(float64(countData) / float64(10)))), TotalData: countData, Attendance: hasil}
		return c.JSON(http.StatusOK, data)
	}
}
