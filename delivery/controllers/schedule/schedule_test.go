package schedule

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
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateSchedule(t *testing.T) {
	t.Run("Failed Create Schedule Because Server Error", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Year":     2222,
			"Month":    2,
			"OfficeId": 1,
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

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.CreateSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 200, res.Code)
		}
	})
	t.Run("Failed Creating Schedule Duplicate Entry", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Year":           2022,
			"Total_capacity": 50,
			"Month":          1,
			"OfficeId":       1,
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

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.CreateSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
	t.Run("Failed Create Schedule Because Wrong Role", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Year":           2022,
			"Total_capacity": 50,
			"Month":          2,
			"OfficeId":       1,
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
		context.SetPath("/schedules")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.CreateSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Create Schedule Because Invalid Token", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Year":     2022,
			"Month":    2,
			"OfficeId": 1,
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
		context.SetPath("/schedules")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.CreateSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Create Schedule Because Failed Binding", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Year":     "asdasd",
			"Month":    2,
			"OfficeId": 1,
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

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.CreateSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Create Schedule Because Server Error", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Year":     2222,
			"Month":    3,
			"OfficeId": 1,
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

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.CreateSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
}

func TestEditSchedule(t *testing.T) {
	t.Run("Success Edit Schedule", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Total_capacity": 100,
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
		context.SetPath("/schedules")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.EditSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Edit Schedule failed binding", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Total_capacity": "100",
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
		context.SetPath("/schedules")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.EditSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Edit Schedule Invalid Token", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Total_capacity": 100,
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
		context.SetPath("/schedules")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.EditSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Edit Schedule Invalid role", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Total_capacity": 100,
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
		context.SetPath("/schedules")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.EditSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Edit Schedule Wrong Param", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Total_capacity": 100,
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
		context.SetPath("/schedules")
		context.SetParamNames("id")
		context.SetParamValues("asdasd")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.EditSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Edit Schedule Internal server error", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Total_capacity": 100,
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
		context.SetPath("/schedules")
		context.SetParamNames("id")
		context.SetParamValues("2")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.EditSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
}

func TestGetScheduleById(t *testing.T) {
	t.Run("Success Get Schedule By Id", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/schedules")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.GetSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Get Schedule By Id Invalid Token", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/schedules")
		context.SetParamNames("id")
		context.SetParamValues("1")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.GetSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Get Schedule By Id Invalid Param", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/schedules")
		context.SetParamNames("id")
		context.SetParamValues("asdasd")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.GetSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Get Schedule By Id Internal Server Error", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/schedules")
		context.SetParamNames("id")
		context.SetParamValues("2")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.GetSchedule("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
}

func TestGetScheduleByMonthAndYear(t *testing.T) {
	t.Run("Success Get Schedule By Month And Year", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/?year=2022&month=2&office=1", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/schedules")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.GetSchedulesByMonthAndYear("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 200, res.Code)
		}
	})
	t.Run("Failed Get Schedule By Month And Year Empty Year", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/?month=2&office=1", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/schedules")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.GetSchedulesByMonthAndYear("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Get Schedule By Month And Year Empty Month", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/?year=2022&office=1", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/schedules")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.GetSchedulesByMonthAndYear("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Get Schedule By Month And Year Empty Office", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/?year=2022&month=2", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/schedules")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.GetSchedulesByMonthAndYear("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Get Schedule By Month And Year Internal Server Error ", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/?year=2022&month=3&office=1", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/schedules")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.GetSchedulesByMonthAndYear("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response common.DefaultResponse
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
	t.Run("Failed Get Schedule By Month And Year Invalid Token", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/?year=2022&month=2&office=1", nil)
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/schedules")

		ScheduleController := New(mockScheduleRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(ScheduleController.GetSchedulesByMonthAndYear("rahasia"))(context)) {
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

type mockScheduleRepository struct{}

func (m mockScheduleRepository) CreateSchedule(month time.Month, year int, capacity int, officeId int) error {
	if month == 3 {
		return fmt.Errorf("error")
	}
	return nil
}

func (m mockScheduleRepository) EditSchedule(scheduleId int, capacity int) error {
	if scheduleId == 2 {
		return fmt.Errorf("failed")
	}
	return nil
}

func (m mockScheduleRepository) GetSchedule(scheduleId int, offset int) (entities.ScheduleResponse, error) {
	if scheduleId == 2 {
		return entities.ScheduleResponse{}, fmt.Errorf("failed")
	}
	return entities.ScheduleResponse{}, nil
}

func (m mockScheduleRepository) GetTotalData(scheduleId int) (int, error) {
	return 1, nil
}

func (m mockScheduleRepository) GetSchedulesByMonthAndYear(month int, year int, officeId int) ([]entities.Schedule, error) {
	if month == 3 {
		return nil, fmt.Errorf("failed")
	}
	if month == 2 {
		return nil, nil
	}
	return []entities.Schedule{}, nil
}
