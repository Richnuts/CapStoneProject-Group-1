package certificate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"sirclo/delivery/common"
	"sirclo/delivery/middlewares"
	"sirclo/entities"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCertificate(t *testing.T) {
	t.Run("Failed Creating Certificate Failed Upload", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		vaccine_dose, _ := w.CreateFormField("vaccine_dose")
		vaccine_dose.Write([]byte("1"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("dosis1"))

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/certificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.CreateCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Creating Certificate Invalid Token", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		vaccine_dose, _ := w.CreateFormField("vaccine_dose")
		vaccine_dose.Write([]byte("1"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("dosis1"))

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(0, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/certificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.CreateCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Creating Certificate Invalid Role", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		vaccine_dose, _ := w.CreateFormField("vaccine_dose")
		vaccine_dose.Write([]byte("1"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("dosis1"))

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/certificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.CreateCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Creating Certificate Already Existed", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		vaccine_dose, _ := w.CreateFormField("vaccine_dose")
		vaccine_dose.Write([]byte("2"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("dosis1"))

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(2, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/certificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.CreateCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusInternalServerError, res.Code)
		}
	})
	t.Run("Failed Creating Certificate Error Binding Image", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		vaccine_dose, _ := w.CreateFormField("vaccine_dose")
		vaccine_dose.Write([]byte("1"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("dosis1"))

		// image, _ := w.CreateFormFile("image", "image.jpg")
		// image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/certificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.CreateCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Creating Certificate Error Invalid Extension", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		vaccine_dose, _ := w.CreateFormField("vaccine_dose")
		vaccine_dose.Write([]byte("1"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("dosis1"))

		image, _ := w.CreateFormFile("image", "image.pdf")
		image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/certificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.CreateCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Creating Certificate Error Invalid Image Size", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		vaccine_dose, _ := w.CreateFormField("vaccine_dose")
		vaccine_dose.Write([]byte("1"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("dosis1"))

		w.CreateFormFile("image", "image.jpg")

		w.Close()
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/certificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.CreateCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Creating Certificate Error Binding Text", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Vaccine_dose": "Approved",
		})
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.CreateCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Creating Certificate Error Vaccine Dose > 3", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Vaccine_dose": 4,
		})
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.CreateCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Creating Certificate Error Dose Exist", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Vaccine_dose": 2,
		})
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.CreateCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
}

func TestGetMyCertificate(t *testing.T) {
	t.Run("Success Get My Certificate", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/mycertificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.GetMyCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Get My Certificate Invalid Token", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/mycertificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.GetMyCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Get My Certificate Invalid Role", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/mycertificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.GetMyCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
}

func TestGetUsersCertificates(t *testing.T) {
	t.Run("Success Get Users Certificates", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.GetUsersCertificates("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Get My Certificate Invalid Token", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.GetUsersCertificates("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Get My Certificate Invalid Role", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.GetUsersCertificates("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
}

func TestGetCertificateById(t *testing.T) {
	t.Run("Success Get Certificate By Id", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.GetCertificateById("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Get My Certificate Invalid Token", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.GetCertificateById("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Get My Certificate Invalid Param", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")
		context.SetParamNames("id")
		context.SetParamValues("satu")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.GetCertificateById("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Get My Certificate Internal Server Error", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(2, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.GetCertificateById("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
}

func TestEditMyCertificate(t *testing.T) {
	t.Run("Failed Edit My Certificate Failed Upload", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/mycertificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditMyCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Edit My Certificate Invalid Token", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(0, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/mycertificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditMyCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Edit My Certificate Invalid Role", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/mycertificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditMyCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Edit My Certificate Status Not Rejected", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(2, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/mycertificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditMyCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusInternalServerError, res.Code)
		}
	})
	t.Run("Failed Edit My Certificate Invalid Param", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/mycertificates")
		context.SetParamNames("id")
		context.SetParamValues("satu")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditMyCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})

	t.Run("Failed Edit My Certificate Dose Exist", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(3, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/mycertificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditMyCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Edit My Certificate Error Binding Image", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		// image, _ := w.CreateFormFile("image", "image.jpg")
		// image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/mycertificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditMyCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Edit My Certificate Error Wrong Extension", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		image, _ := w.CreateFormFile("image", "image.pdf")
		image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/mycertificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditMyCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Edit My Certificate Error Wrong Extension", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		w.CreateFormFile("image", "image.jpg")

		w.Close()
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/mycertificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditMyCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
}

func TestEditCertificate(t *testing.T) {
	t.Run("Success Edit Certificate Status", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Status": "Approved",
		})
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Edit Certificate Status Invalid Token", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Status": "Approved",
		})
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Edit Certificate Status Invalid Role", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Status": "Approved",
		})
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Edit Certificate Status Invalid Param", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Status": "Approved",
		})
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")
		context.SetParamNames("id")
		context.SetParamValues("satu")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Edit Certificate Status Error Binding Text", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Status": 1,
		})
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Edit Certificate Status Invalid Status", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Status": "Verified",
		})
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Edit Certificate Status Internal Server Error", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Status": "Approved",
		})
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/certificates")
		context.SetParamNames("id")
		context.SetParamValues("2")

		CertificateController := New(mockCertificateRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CertificateController.EditCertificate("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusInternalServerError, res.Code)
		}
	})
}

// =========================== mocking ===========================

type mockCertificateRepository struct{}

func (m mockCertificateRepository) CreateCertificate(userId, vaccineDose int, imageURL, description string) error {
	if vaccineDose == 4 {
		return fmt.Errorf("jangan ngaco")
	}
	return nil
}

func (m mockCertificateRepository) GetMyCertificate(userId int) (entities.UsersCertificateWithName, error) {
	var hasil entities.UsersCertificateWithName
	return hasil, nil
}

func (m mockCertificateRepository) GetUsersCertificates(status string, offset int) ([]entities.UsersCertificateWithName, error) {
	var hasil []entities.UsersCertificateWithName
	return hasil, nil
}

func (m mockCertificateRepository) GetCertificateById(id, userId int) (entities.CertificateResponseGetByIdAndUID, error) {
	var hasil entities.CertificateResponseGetByIdAndUID
	if userId == 2 {
		return hasil, fmt.Errorf("internal server error")
	}
	return hasil, nil
}

func (m mockCertificateRepository) EditCertificate(id, adminId int, status string) error {
	if id == 2 {
		return fmt.Errorf("error")
	}
	return nil
}

func (m mockCertificateRepository) EditMyCertificate(id int, imageURL string) error {
	return nil
}

func (m mockCertificateRepository) GetCertificateByDose(userId, vaccineDose int) error {
	if userId == 3 {
		return fmt.Errorf("error")
	}
	if vaccineDose == 2 {
		return fmt.Errorf("error")
	}
	return nil
}

func (m mockCertificateRepository) GetVaccineStatus(userId int) error {
	if userId == 2 {
		return fmt.Errorf("error")
	}
	return nil
}

func (m mockCertificateRepository) GetVaccineDose(id int) (int, error) {
	return 0, nil
}

func (m mockCertificateRepository) GetTotalPage(status string) (int, error) {
	return 0, nil
}

func (m mockCertificateRepository) GetTotalUsers() (int, error) {
	return 0, nil
}
