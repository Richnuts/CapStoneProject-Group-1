package office

import (
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

func TestGetOffice(t *testing.T) {
	t.Run("Success Get Office By Id", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		fmt.Println(token)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/offices")
		context.SetParamNames("id")
		context.SetParamValues("1")

		OfficeController := New(mockOfficeRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(OfficeController.GetOffice("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			fmt.Println(response)
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Get Office By Id Wrong token", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		fmt.Println(token)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/offices")
		context.SetParamNames("id")
		context.SetParamValues("1")

		OfficeController := New(mockOfficeRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(OfficeController.GetOffice("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			fmt.Println(response)
			assert.Equal(t, 403, res.Code)
		}

	})
	t.Run("Failed Get Office wrong param", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		fmt.Println(token)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/offices")
		context.SetParamNames("id")
		context.SetParamValues("asd")

		OfficeController := New(mockOfficeRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(OfficeController.GetOffice("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			fmt.Println(response)
			assert.Equal(t, 400, res.Code)
		}
	})
}

func TestGetOffices(t *testing.T) {
	t.Run("Success Get Offices", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		fmt.Println(token)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/offices")

		OfficeController := New(mockOfficeRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(OfficeController.GetOffices("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			fmt.Println(response)
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Get Offices Wrong token", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(0, "admin")
		if err != nil {
			panic(err)
		}
		fmt.Println(token)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/offices")

		OfficeController := New(mockOfficeRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(OfficeController.GetOffices("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			fmt.Println(response)
			assert.Equal(t, 403, res.Code)
		}
	})
}

type mockOfficeRepository struct{}

func (m mockOfficeRepository) GetOffice(officeId int) (entities.Office, error) {
	return entities.Office{
		Id:     1,
		Name:   "asd",
		Detail: "asd",
	}, nil
}

func (m mockOfficeRepository) GetOffices() ([]entities.Office, error) {
	return []entities.Office{{
		Id:     1,
		Name:   "asd",
		Detail: "asd",
	},
		{
			Id:     1,
			Name:   "asd",
			Detail: "asd",
		}}, nil
}
