package attendance

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
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateAttendance(t *testing.T) {
	t.Run("Failed Creating Attendance Failed Upload", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		schedule_id, _ := w.CreateFormField("schedule_id")
		schedule_id.Write([]byte("1"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("pcr"))

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("string apapun"))

		w.Close()
		token, err := middlewares.CreateToken(777, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/schedules")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.CreateAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
	t.Run("Failed Creating Attendance Wrong image size", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		schedule_id, _ := w.CreateFormField("schedule_id")
		schedule_id.Write([]byte("1"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("pcr"))

		w.CreateFormFile("image", "image.jpg")

		w.Close()
		token, err := middlewares.CreateToken(777, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/schedules")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.CreateAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Creating Attendance Wrong image ext", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		schedule_id, _ := w.CreateFormField("schedule_id")
		schedule_id.Write([]byte("1"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("pcr"))

		w.CreateFormFile("image", "image.pdf")

		w.Close()
		token, err := middlewares.CreateToken(777, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/schedules")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.CreateAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Creating Attendance Binding Failed", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		schedule_id, _ := w.CreateFormField("schedule_id")
		schedule_id.Write([]byte("1"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("pcr"))

		w.Close()
		token, err := middlewares.CreateToken(777, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/schedules")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.CreateAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Creating Attendance Request Already Exist", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		schedule_id, _ := w.CreateFormField("schedule_id")
		schedule_id.Write([]byte("1"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("pcr"))

		w.CreateFormFile("image", "image.jpg")

		w.Close()
		token, err := middlewares.CreateToken(2, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/schedules")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.CreateAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Creating Attendance Request Time < Time.Now", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		schedule_id, _ := w.CreateFormField("schedule_id")
		schedule_id.Write([]byte("3"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("pcr"))

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("pcr"))

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
		context.SetPath("/schedules")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.CreateAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Creating Attendance Request Time Not Found in DB", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		schedule_id, _ := w.CreateFormField("schedule_id")
		schedule_id.Write([]byte("2"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("pcr"))

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("pcr"))

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
		context.SetPath("/schedules")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.CreateAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Creating Attendance User Unhealty", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		schedule_id, _ := w.CreateFormField("schedule_id")
		schedule_id.Write([]byte("1"))

		description, _ := w.CreateFormField("description")
		description.Write([]byte("pcr"))

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("pcr"))

		w.Close()
		token, err := middlewares.CreateToken(3, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/schedules")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.CreateAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Creating Attendance Failed Binding", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"Schedule_id": "asdqwe",
			"Description": "asd",
		})
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/schedules")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.CreateAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Creating Attendance Invalid Token", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		image, _ := w.CreateFormFile("image", "image.jpg")
		image.Write([]byte("pcr"))

		w.Close()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", buf)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		context := e.NewContext(req, res)
		context.SetPath("/schedules")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.CreateAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
}

func TestEditAttendance(t *testing.T) {
	t.Run("Success Edit Attendance", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Attendance_id": 29,
			"Status":        "Approved",
			"Status_info":   "boleh kosong kok",
		})
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/attendances")
		context.SetParamNames("id")
		context.SetParamValues("1")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.EditAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 200, res.Code)
		}
	})
	t.Run("Failed Edit Attendance Invalid Token", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Attendance_id": 29,
			"Status":        "Approved",
			"Status_info":   "boleh kosong kok",
		})
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.EditAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Edit Attendance Invalid Role", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Attendance_id": 29,
			"Status":        "Approved",
			"Status_info":   "boleh kosong kok",
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
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.EditAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Edit Attendance Invalid Param", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Attendance_id": 29,
			"Status":        "Approved",
			"Status_info":   "boleh kosong kok",
		})
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("asdasd")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.EditAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Edit Attendance Binding Failed", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Schedule_id": "asdasd",
			"Status":      "Approved",
			"Status_info": "boleh kosong kok",
		})
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.EditAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Edit Attendance Check Capacity Error", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Schedule_id": 20,
			"Status":      "Approved",
		})
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/attendances")
		context.SetParamNames("id")
		context.SetParamValues("20")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.EditAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
	t.Run("Failed Edit Attendance Check Capacity Error", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Schedule_id": 22,
			"Status":      "Approved",
		})
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/attendances")
		context.SetParamNames("id")
		context.SetParamValues("22")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.EditAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
	t.Run("Failed Edit Attendance Internal Server Error", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Attendance_id": 19,
			"Status":        "Approved",
			"Status_info":   "boleh kosong kok",
		})
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/attendances")
		context.SetParamNames("id")
		context.SetParamValues("3")

		AttendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(AttendanceController.EditAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
}

func TestGetAttendanceById(t *testing.T) {
	t.Run("Success Get Attendance By ID", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/attendances")
		context.SetParamNames("id")
		context.SetParamValues("1")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetAttendanceById("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Success Get Attendance By ID Invalid Token", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/attendances")
		context.SetParamNames("id")
		context.SetParamValues("1")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetAttendanceById("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Success Get Attendance By ID Invalid Param", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/attendances")
		context.SetParamNames("id")
		context.SetParamValues("asdasd")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetAttendanceById("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Success Get Attendance By ID Internal Server Error", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/attendances")
		context.SetParamNames("id")
		context.SetParamValues("2")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetAttendanceById("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
}

func TestGetMyAttendance(t *testing.T) {
	t.Run("Success Get My Attendance", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/attendances")
		context.SetParamNames("id")
		context.SetParamValues("1")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetMyAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Get My Attendance Internal Server Error", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(3, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/attendances")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetMyAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
	t.Run("Failed Get My Attendance Invalid token", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/attendances")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetMyAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
}

func TestGetMyLatestAttendance(t *testing.T) {
	t.Run("Success Get My Attendance By Lates", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/attendances")
		context.SetParamNames("id")
		context.SetParamValues("1")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetMyAttendanceSortByLatest("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Get My Attendance By Latest Internal Server Error", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(3, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/attendances")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetMyAttendanceSortByLatest("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
	t.Run("Failed Get My Attendance By Latest Invalid token", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/attendances")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetMyAttendanceSortByLatest("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
}

func TestGetMyLongestAttendance(t *testing.T) {
	t.Run("Success Get My Attendance By Longest", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/attendances")
		context.SetParamNames("id")
		context.SetParamValues("1")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetMyAttendanceSortByLongest("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Get My Attendance By Longest Internal Server Error", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(3, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/attendances")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetMyAttendanceSortByLongest("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
	t.Run("Failed Get My Attendance By Longest Invalid token", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/attendances")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetMyAttendanceSortByLongest("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
}

func TestGetPendingAttendance(t *testing.T) {
	t.Run("Success Get Pending Attendance", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/?office=1", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/pendingattendances")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetPendingAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Get Pending Attendance Internal Server Error", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(3, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/?office=2", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/pendingattendances")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetPendingAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
	t.Run("Failed Get Pending Attendance Invalid token", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/?office=1", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/pendingattendances")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetPendingAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Get Pending Attendance Invalid Role", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/?office=1", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/pendingattendances")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetPendingAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Get Pending Attendance Invalid office", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/pendingattendances")

		attendanceController := New(mockAttendanceRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(attendanceController.GetPendingAttendance("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
}

type mockAttendanceRepository struct{}

func (m mockAttendanceRepository) CreateAttendance(userId int, scheduleId int, description string, imageUrl string) error {
	return nil
}

func (m mockAttendanceRepository) EditAttendance(attendanceId int, adminId int, status string, statusInfo string) error {
	if attendanceId == 3 {
		return fmt.Errorf("error")
	}
	return nil
}

func (m mockAttendanceRepository) GetPendingAttendance(offset int, officeId int) ([]entities.PendingAttendance, error) {
	if officeId == 2 {
		return []entities.PendingAttendance{}, fmt.Errorf("error")
	}
	return []entities.PendingAttendance{}, nil
}

func (m mockAttendanceRepository) GetPendingAttendanceTotalData(officeId int) (int, error) {
	return 1, nil
}

func (m mockAttendanceRepository) GetMyAttendance(userId int, offset int, status string) ([]entities.AttendanceGetFormat, error) {
	if userId == 3 {
		return []entities.AttendanceGetFormat{}, fmt.Errorf("error")
	}
	return []entities.AttendanceGetFormat{}, nil
}

func (m mockAttendanceRepository) GetAttendanceById(attendanceId int) (entities.AttendanceGetFormat, error) {
	if attendanceId == 2 {
		return entities.AttendanceGetFormat{}, fmt.Errorf("failed")
	}
	return entities.AttendanceGetFormat{}, nil
}

func (m mockAttendanceRepository) GetMyAttendanceSortByLatest(userId int, offset int, status string) ([]entities.AttendanceGetFormat, error) {
	if userId == 3 {
		return []entities.AttendanceGetFormat{}, fmt.Errorf("error")
	}
	return []entities.AttendanceGetFormat{}, nil
}

func (m mockAttendanceRepository) GetMyAttendanceSortByLongest(userId int, offset int, status string) ([]entities.AttendanceGetFormat, error) {
	if userId == 3 {
		return []entities.AttendanceGetFormat{}, fmt.Errorf("error")
	}
	return []entities.AttendanceGetFormat{}, nil
}

func (m mockAttendanceRepository) GetMyAttendanceTotalData(userId int, status string) (int, error) {
	return 1, nil
}

func (m mockAttendanceRepository) GetUserAttendanceStatus(userId int, scheduleId int) error {
	if userId == 2 {
		return fmt.Errorf("request sudah ada")
	}
	return nil
}

func (m mockAttendanceRepository) GetUserVaccineStatus(userId int) error {
	if userId == 3 {
		return fmt.Errorf("unhealthy")
	}
	return nil
}

func (m mockAttendanceRepository) CheckCapacity(scheduleId int) (int, error) {
	if scheduleId == 20 {
		return 1, fmt.Errorf("capacity habis")
	}
	if scheduleId == 22 {
		return 0, nil
	}
	return 1, nil
}

func (m mockAttendanceRepository) CheckCreateRequestDate(scheduleId int) (time.Time, error) {
	if scheduleId == 3 {
		return time.Now().AddDate(-1, -1, -1), nil
	}
	if scheduleId == 2 {
		return time.Now(), fmt.Errorf("error time not found")
	}
	return time.Now().AddDate(1, 1, 1), nil
}
