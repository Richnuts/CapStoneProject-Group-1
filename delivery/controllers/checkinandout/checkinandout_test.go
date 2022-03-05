package checkinandout

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sirclo/delivery/common"
	"sirclo/delivery/middlewares"
	"sirclo/entities"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCheckin(t *testing.T) {
	t.Run("Success Checkin", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Id":          1,
			"Temperature": 36.0,
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
		context.SetPath("/checkin")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.Checkin("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Success Checkin But Rejected", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Id":          1,
			"Temperature": 38.0,
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
		context.SetPath("/checkin")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.Checkin("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Checkin Invalid Token", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Id":          1,
			"Temperature": 36.0,
		})
		token, err := middlewares.CreateToken(0, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/checkin")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.Checkin("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Checkin Invalid Role", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Id":          1,
			"Temperature": 36.0,
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
		context.SetPath("/checkin")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.Checkin("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Checkin Binding Failed", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Id":          "satu",
			"Temperature": 36,
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
		context.SetPath("/checkin")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.Checkin("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Checkin Wrong Date", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Id":          2,
			"Temperature": 36.0,
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
		context.SetPath("/checkin")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.Checkin("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Checkin Internal Server Error", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Id":          1,
			"Temperature": 36.0,
		})
		token, err := middlewares.CreateToken(2, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/checkin")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.Checkin("rahasia"))(context)) {
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

func TestCheckout(t *testing.T) {
	t.Run("Success Checkout", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Id": 1,
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
		context.SetPath("/checkout")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.Checkout("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Checkout Invalid Token", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Id": 1,
		})
		token, err := middlewares.CreateToken(0, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/checkout")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.Checkout("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Checkout Invalid Role", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Id": 1,
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
		context.SetPath("/checkout")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.Checkout("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Checkout Binding Failed", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Id": "satu",
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
		context.SetPath("/checkout")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.Checkout("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Failed Checkout Internal Server Error", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Id": 1,
		})
		token, err := middlewares.CreateToken(2, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/checkout")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.Checkout("rahasia"))(context)) {
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

func TestGetCheckById(t *testing.T) {
	t.Run("Success Get Check By Id", func(t *testing.T) {
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
		context.SetPath("/checks")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.GetCheckById("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Error Get Check By Id Invalid Token", func(t *testing.T) {
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
		context.SetPath("/checks")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.GetCheckById("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Error Get Check By Id Invalid Param", func(t *testing.T) {
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
		context.SetPath("/checks")
		context.SetParamNames("id")
		context.SetParamValues("asdasd")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.GetCheckById("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
	t.Run("Error Get Check By Id Id Not Found", func(t *testing.T) {
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
		context.SetPath("/checks")
		context.SetParamNames("id")
		context.SetParamValues("2")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.GetCheckById("rahasia"))(context)) {
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

func TestGetAllCheck(t *testing.T) {
	t.Run("Success Get All Check", func(t *testing.T) {
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
		context.SetPath("/checkschedule")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.GetAllCheck("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Get All Check Invalid Token", func(t *testing.T) {
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
		context.SetPath("/checkschedule")
		context.SetParamNames("id")
		context.SetParamValues("1")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.GetAllCheck("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Get All Check Invalid Param", func(t *testing.T) {
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
		context.SetPath("/checkschedule")
		context.SetParamNames("id")
		context.SetParamValues("satu")

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.GetAllCheck("rahasia"))(context)) {
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

// =========================== mocking ===========================

type mockCheckRepository struct{}

func (m mockCheckRepository) Checkin(attendanceId, userId int, temperature float64, status string) error {
	if attendanceId == 2 {
		return fmt.Errorf("attendance not found")
	}
	if userId == 2 {
		return fmt.Errorf("wrong user")
	}
	return nil
}

func (m mockCheckRepository) Checkout(attendanceId, userId int) error {
	if attendanceId == 2 {
		return fmt.Errorf("attendance not found")
	}
	if userId == 2 {
		return fmt.Errorf("wrong user")
	}
	return nil
}

// func (m mockCheckRepository) GetAllCheck(offset int) ([]entities.GetAllCheck, error) {
// 	var hasil []entities.GetAllCheck
// 	return hasil, nil
// }

func (m mockCheckRepository) GetAllCheck(id, offset int) ([]entities.GetCheckbyDate, error) {
	var hasil []entities.GetCheckbyDate
	return hasil, nil
}

func (m mockCheckRepository) GetCheckDate(id int) error {
	if id == 2 {
		return fmt.Errorf("id not found")
	}
	return nil
}

func (m mockCheckRepository) GetCheckbyId(id int) (entities.CheckinAndOutResponseFormat, error) {
	var hasil entities.CheckinAndOutResponseFormat
	if id == 2 {
		return hasil, fmt.Errorf("error")
	}
	return hasil, nil
}

func (m mockCheckRepository) GetTotalPage() (int, error) {

	return 1, nil
}
