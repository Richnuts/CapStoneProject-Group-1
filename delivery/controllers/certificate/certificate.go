package certificate

import (
	"fmt"
	"net/http"
	"sirclo/delivery/common"
	"sirclo/delivery/controllers/imageLib"
	"sirclo/delivery/middlewares"
	"sirclo/entities"
	certificateRepo "sirclo/repository/certificate"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type CertificateController struct {
	repository certificateRepo.Certificate
}

func New(certificate certificateRepo.Certificate) *CertificateController {
	return &CertificateController{
		repository: certificate,
	}
}

func (cer CertificateController) CreateCertificate(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		var certificateRequest CertificateRequestFormat
		// prosess binding text
		if err_bind := c.Bind(&certificateRequest); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		err_checking := cer.repository.GetCertificateByDose(loginId, certificateRequest.VaccineDose)
		if err_checking != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "operation failed", "request telah ada"))
		}
		// prosess binding image
		fileData, fileInfo, err_binding_image := c.Request().FormFile("image")
		if err_binding_image != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "bad request", "bind image error"))
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
		fileName := "certificate_" + strconv.Itoa(loginId) + "_" + strconv.Itoa(certificateRequest.VaccineDose)

		// upload the photo
		var err_upload_photo error
		theUrl, err_upload_photo := imageLib.UploadImage("certificate", fileName, fileData)
		if err_upload_photo != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "bad request", "Upload Image Failed"))
		}
		// create certificate
		imageURL := theUrl
		err_certificate := cer.repository.CreateCertificate(loginId, certificateRequest.VaccineDose, imageURL, certificateRequest.Description)
		if err_certificate != nil {
			fmt.Println(err_certificate)
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "failed creating schedule", "duplicate entry"))
		}
		return c.JSON(http.StatusOK, common.SuccessOperation("berhasil membuat jadwal"))
	}
}

func (cer CertificateController) GetCertificate(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// mengGet certificate
		var data []entities.CertificateResponseGetByIdAndUID
		data, err_get := cer.repository.GetApprovedCertificate(loginId)
		if err_get != nil {
			return c.JSON(http.StatusBadRequest, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, data)
	}
}

func (cer CertificateController) GetCertificateById(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// getting the id
		certificateId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		// mengGet certificate
		var data entities.CertificateResponseGetByIdAndUID
		data, err_get := cer.repository.GetCertificateById(certificateId, loginId)
		if err_get != nil {
			return c.JSON(http.StatusBadRequest, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, data)
	}
}

func (cer CertificateController) EditCertificate(secret string) echo.HandlerFunc {
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
		certificateId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		var CertificateEditRequest CertificateEditFormat
		// prosess binding text
		if err_bind := c.Bind(&CertificateEditRequest); err_bind != nil {
			fmt.Println(err_bind)
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		err_edit := cer.repository.EditCertificate(certificateId, CertificateEditRequest.Status)

		if err_edit != nil {
			fmt.Println(err_edit)
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation("check out success"))
	}
}
