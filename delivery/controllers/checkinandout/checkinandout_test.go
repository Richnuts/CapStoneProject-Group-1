package checkinandout

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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
			"Id":          "1",
			"Temperature": "36.0",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		CheckController := New(mockCheckRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(CheckController.Checkin("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			// var response entities.Login
			err := json.Unmarshal([]byte(bodyResponses), "success")
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
}

// =========================== mocking ===========================

type mockCheckRepository struct{}

func (m mockCheckRepository) Checkin(attendanceId, userId int, temperature float64, status string) error {
	if attendanceId == 2 {
		return fmt.Errorf("attendance not found")
	}
	return nil
}

func (m mockCheckRepository) Checkout(attendanceId, userId int) error {
	if attendanceId == 2 {
		return fmt.Errorf("attendance not found")
	}
	return nil
}

func (m mockCheckRepository) GetAllCheck(offset int) ([]entities.GetAllCheck, error) {
	var hasil []entities.GetAllCheck
	return hasil, nil
}

func (m mockCheckRepository) GetCheckDate(id int) error {

	return nil
}

func (m mockCheckRepository) GetCheckbyId(id int) (entities.CheckinAndOutResponseFormat, error) {
	var hasil entities.CheckinAndOutResponseFormat
	return hasil, nil
}

func (m mockCheckRepository) GetTotalPage() (int, error) {

	return 1, nil
}
