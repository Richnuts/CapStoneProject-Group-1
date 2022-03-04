package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"sirclo/delivery/middlewares"
	"sirclo/entities"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	t.Run("Success Get User By Id", func(t *testing.T) {
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
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Get User By Wrong Token", func(t *testing.T) {
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
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusUnauthorized, res.Code)
		}
	})
	t.Run("Failed Get User By Wrong Role", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Get User By internal server error", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(2, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("2")

		userController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
	t.Run("Failed Get User By Wrong Param", func(t *testing.T) {
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
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("asdasd")

		userController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
}
func TestGetProfile(t *testing.T) {
	t.Run("Success Get Profile", func(t *testing.T) {
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
		context.SetPath("/users")

		userController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetProfile("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Get Profile Invalid token", func(t *testing.T) {
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
		context.SetPath("/users")

		userController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetProfile("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusUnauthorized, res.Code)
		}
	})
	t.Run("Failed Get Profile Internal server error", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(2, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.GetProfile("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusInternalServerError, res.Code)
		}
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Success Delete User", func(t *testing.T) {
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
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.DeleteUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
	t.Run("Failed Delete User By Wrong Token", func(t *testing.T) {
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
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.DeleteUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusUnauthorized, res.Code)
		}
	})
	t.Run("Failed Delete User Token and id not match", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "user")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("2")

		userController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.DeleteUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusForbidden, res.Code)
		}
	})
	t.Run("Failed Get User By internal server error", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(2, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("2")

		userController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.DeleteUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
	t.Run("Failed Get User By Wrong Param", func(t *testing.T) {
		e := echo.New()
		token, err := middlewares.CreateToken(1, "admin")
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodGet, "/?p=2&rp=2", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("asdasd")
		userController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(userController.DeleteUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.Office
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, res.Code)
		}
	})
}

func TestEditUser(t *testing.T) {
	t.Run("Success Edit Profile", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"Name":     "asda",
			"Password": "asd",
			"Email":    "asd",
			"ImageUrl": "asd",
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

		UserController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(UserController.EditUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 200, res.Code)
		}
	})
	t.Run("Failed Edit Profile Wrong Token", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"Name":     "asda",
			"Password": "asd",
			"Email":    "asd",
			"ImageUrl": "asd",
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
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		UserController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(UserController.EditUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 401, res.Code)
		}
	})
	t.Run("Failed Edit Profile Wrong Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"Name":     123,
			"Password": "asd",
			"Email":    "asd",
			"ImageUrl": "asd",
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

		UserController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(UserController.EditUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Edit Profile Token and id not match", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"Name":     "asda",
			"Password": "asd",
			"Email":    "asd",
			"ImageUrl": "asd",
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
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		UserController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(UserController.EditUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 403, res.Code)
		}
	})
	t.Run("Failed Edit Profile Image Url not found", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"Name":     "asda",
			"Password": "asd",
			"Email":    "asd",
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
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("2")

		UserController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(UserController.EditUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
	t.Run("Failed Edit Profile internal server error", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"Name":     "asda",
			"Password": "asd",
			"Email":    "asd",
			"ImageUrl": "asd",
		})
		token, err := middlewares.CreateToken(3, "user")
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
		context.SetParamValues("3")

		UserController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(UserController.EditUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 500, res.Code)
		}
	})
	t.Run("Failed Edit Profile Wrong Param", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"Name":     "asda",
			"Password": "asd",
			"Email":    "asd",
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
		context.SetParamValues("asdasd")

		UserController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(UserController.EditUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Edit Profile Wrong Image Ext", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		name, _ := w.CreateFormField("name")
		name.Write([]byte("bangkus"))

		email, _ := w.CreateFormField("email")
		email.Write([]byte("bangkus@gmail.com"))

		password, _ := w.CreateFormField("password")
		password.Write([]byte("asdasd"))

		w.CreateFormFile("image", "image.pdf")

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
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		UserController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(UserController.EditUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Edit Profile Wrong Image Size", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		name, _ := w.CreateFormField("name")
		name.Write([]byte("bangkus"))

		email, _ := w.CreateFormField("email")
		email.Write([]byte("bangkus@gmail.com"))

		password, _ := w.CreateFormField("password")
		password.Write([]byte("asdasd"))

		w.CreateFormFile("image", "image.png")

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
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		UserController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(UserController.EditUser("rahasia"))(context)) {
			bodyResponses := res.Body.String()
			var response entities.UserResponseFormat
			err := json.Unmarshal([]byte(bodyResponses), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, 400, res.Code)
		}
	})
	t.Run("Failed Edit Profile Failed Upload", func(t *testing.T) {
		e := echo.New()
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		name, _ := w.CreateFormField("name")
		name.Write([]byte("bangkus"))

		email, _ := w.CreateFormField("email")
		email.Write([]byte("bangkus@gmail.com"))

		password, _ := w.CreateFormField("password")
		password.Write([]byte("asdasd"))

		image, _ := w.CreateFormFile("image", "image.png")
		image.Write([]byte("asdasd"))

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
		context.SetPath("/users")
		context.SetParamNames("id")
		context.SetParamValues("1")

		UserController := New(mockUserRepository{})
		if assert.NoError(t, middlewares.JWTMiddleware()(UserController.EditUser("rahasia"))(context)) {
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

type mockUserRepository struct{}

func (m mockUserRepository) GetUser(userId int) (entities.UserResponseFormat, error) {
	if userId == 2 {
		return entities.UserResponseFormat{}, fmt.Errorf("error")
	}
	return entities.UserResponseFormat{}, nil
}

func (m mockUserRepository) EditUser(user entities.User) error {
	if user.Id == 3 {
		return fmt.Errorf("error")
	}
	return nil
}

func (m mockUserRepository) DeleteUser(id int) error {
	if id == 2 {
		return fmt.Errorf("error")
	}
	return nil
}

func (m mockUserRepository) GetUserImageUrl(id int) string {
	if id == 2 {
		return ""
	}
	return "www.fotoku.com"
}
