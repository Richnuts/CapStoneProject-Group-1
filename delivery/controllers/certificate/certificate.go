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
		// check role
		role := middlewares.GetUserRole(secret, c)
		fmt.Println(role)
		if role != "user" {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// check status vaccine user
		err_status := cer.repository.GetVaccineStatus(loginId)
		if err_status != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		var certificateRequest CertificateRequestFormat
		// prosess binding text
		if err_bind := c.Bind(&certificateRequest); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		err_checking := cer.repository.GetCertificateByDose(loginId, certificateRequest.VaccineDose)
		if certificateRequest.VaccineDose > 3 {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "operation failed", "jangan ngaco"))
		}
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
		return c.JSON(http.StatusOK, common.SuccessOperation("berhasil menambah sertifikat"))
	}
}

func (cer CertificateController) GetMyCertificate(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// check role
		role := middlewares.GetUserRole(secret, c)
		fmt.Println(role)
		if role != "user" {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// mengGet certificate
		var data entities.UsersCertificateWithName
		data, err_get := cer.repository.GetMyCertificate(loginId)
		if err_get != nil {
			fmt.Println(err_get)
			return c.JSON(http.StatusBadRequest, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, data)
	}
}

func (cer CertificateController) GetUsersCertificates(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// check role
		role := middlewares.GetUserRole(secret, c)
		fmt.Println(role)
		if role != "admin" {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// getting the status
		status := c.QueryParam("status")
		// getting the page
		pageString := c.QueryParam("page")
		halaman, err := strconv.Atoi(pageString)
		if err != nil {
			halaman = 1
		}
		fmt.Println("halamannya = ", halaman)
		offset := (halaman - 1) * 10
		// mengGet list
		var hasil entities.UsersCertificateWithPage
		// var name entities.UsersCertificate
		var err_get error
		// hasil.Name, _ = cer.repository.GetName(name.Id)
		hasil.Certificates, err_get = cer.repository.GetUsersCertificates(status, offset)
		if err_get != nil {
			fmt.Println("anu", err_get)
			return c.JSON(http.StatusBadRequest, common.InternalServerError())
		}
		var err_totalusers error
		hasil.TotalUsers, err_totalusers = cer.repository.GetTotalUsers()
		if err_totalusers != nil {
			return c.JSON(http.StatusBadRequest, common.InternalServerError())
		}
		// menGet total page
		var err_page error
		hasil.TotalPage, err_page = cer.repository.GetTotalPage(status)
		if err_page != nil {
			return c.JSON(http.StatusBadRequest, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, hasil)
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

func (cer CertificateController) EditMyCertificate(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		fmt.Println(loginId)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// check role
		role := middlewares.GetUserRole(secret, c)
		fmt.Println(role)
		if role != "user" {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// check status vaccine user
		err_status := cer.repository.GetVaccineStatus(loginId)
		if err_status != nil {
			fmt.Println("1", err_status)
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		// getting the id
		certificateId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println("2", err)
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		var CertificateEditRequest CertificateEditFormat
		// prosess binding text
		if err_bind := c.Bind(&CertificateEditRequest); err_bind != nil {
			fmt.Println(err_bind)
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		vaccineDose, _ := cer.repository.GetVaccineDose(certificateId)
		fmt.Println("dosis", vaccineDose)
		err_checking := cer.repository.GetCertificateByDose(loginId, vaccineDose)
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
		fileName := "certificate_" + strconv.Itoa(loginId) + "_" + strconv.Itoa(vaccineDose)

		// upload the photo
		var err_upload_photo error
		theUrl, err_upload_photo := imageLib.UploadImage("certificate", fileName, fileData)
		if err_upload_photo != nil {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "bad request", "Upload Image Failed"))
		}
		imageURL := theUrl
		err_edit := cer.repository.EditMyCertificate(certificateId, imageURL)

		if err_edit != nil {
			fmt.Println(err_edit)
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation("Edit image success"))
	}
}

func (cer CertificateController) EditCertificate(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		fmt.Println(loginId)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// check role
		role := middlewares.GetUserRole(secret, c)
		fmt.Println(role)
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
		if CertificateEditRequest.Status != "Rejected" && CertificateEditRequest.Status != "Approved" {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		err_edit := cer.repository.EditCertificate(certificateId, loginId, CertificateEditRequest.Status)

		if err_edit != nil {
			fmt.Println(err_edit)
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation("Edit status success"))
	}
}
