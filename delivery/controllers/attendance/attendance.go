package attendance

import (
	"fmt"
	"net/http"
	"sirclo/delivery/common"
	"sirclo/delivery/controllers/imageLib"
	"sirclo/delivery/middlewares"
	attendanceRepo "sirclo/repository/attendance"
	"strconv"

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
		var attendanceRequest AttendanceRequestFormat
		// prosess binding text
		if err_bind := c.Bind(&attendanceRequest); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "operation failed", "binding text gagal"))
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
			fmt.Println(err_create)
			return c.JSON(http.StatusBadRequest, common.CustomResponse(500, "internal server error", "Failed Creating Entity in database"))
		}
		fmt.Println(theUrl)
		return c.JSON(http.StatusOK, common.CustomResponse(200, "operation success", "berhasil membuat request WFO"))
	}
}
