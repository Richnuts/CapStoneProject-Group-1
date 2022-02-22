package office

import (
	"net/http"
	"sirclo/delivery/common"
	"sirclo/delivery/middlewares"
	officeRepo "sirclo/repository/office"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OfficeController struct {
	repository officeRepo.Office
}

func New(office officeRepo.Office) *OfficeController {
	return &OfficeController{
		repository: office,
	}
}

func (oc OfficeController) GetOffices(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// getting offices
		data, _ := oc.repository.GetOffices()
		return c.JSON(http.StatusOK, data)
	}
}

func (oc OfficeController) GetOffice(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// getting the id
		officeId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		// getting office
		data, _ := oc.repository.GetOffice(officeId)
		return c.JSON(http.StatusOK, data)
	}
}
