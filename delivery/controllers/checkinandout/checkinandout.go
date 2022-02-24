package checkinandout

import (
	"net/http"
	"sirclo/delivery/common"
	"sirclo/delivery/middlewares"
	checkRepo "sirclo/repository/checkinandout"

	echo "github.com/labstack/echo/v4"
)

type CheckController struct {
	repository checkRepo.CheckinAndOut
}

func New(attendance checkRepo.CheckinAndOut) *CheckController {
	return &CheckController{
		repository: attendance,
	}
}

func (cc CheckController) Checkin(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		var CheckinRequest CheckinRequestFormat
		status := "Approved"
		// prosess binding text
		if err_bind := c.Bind(&CheckinRequest); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		if CheckinRequest.Temperature >= 37.5 {
			status = "Rejected"
		}
		err_edit := cc.repository.Checkin(CheckinRequest.Id, loginId, CheckinRequest.Temperature, status)
		if err_edit != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation("check in success"))
	}
}

func (cc CheckController) Checkout(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		var CheckoutRequest CheckoutRequestFormat
		// prosess binding text
		if err_bind := c.Bind(&CheckoutRequest); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		err_edit := cc.repository.Checkout(CheckoutRequest.Id, loginId)

		if err_edit != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation("check out success"))
	}
}
