package auth

import (
	"net/http"
	"sirclo/delivery/common"
	authRepo "sirclo/repository/auth"

	echo "github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	repository authRepo.Auth
}

func New(auth authRepo.Auth) *AuthController {
	return &AuthController{
		repository: auth,
	}
}

func (a AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var loginRequest LoginRequestFormat
		// binding data mendapatkan loginRequest
		if err := c.Bind(&loginRequest); err != nil {
			return c.JSON(http.StatusForbidden, common.CustomResponse(400, "Failed Binding", "Failed to bind JSON input"))
		}
		// mengambil data user dari DB berdasarkan EMAIL
		user, err_login := a.repository.FindUserByEmail(loginRequest.Email)
		if err_login != nil {
			return c.JSON(http.StatusForbidden, common.CustomResponse(403, "Failed Checking Email", "Email Not Match"))
		}
		// check password
		err_token := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
		if err_token != nil {
			return c.JSON(http.StatusForbidden, common.CustomResponse(403, "Failed Checking Password", "Password Not Match"))
		}

		loginResponse, err := a.repository.Login(loginRequest.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, loginResponse)
	}
}

// func (a AuthController) Register() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var userRequest UserRequestFormat

// 		if err := c.Bind(&userRequest); err != nil {
// 			return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "Failed Binding", "Failed to bind JSON input"))
// 		}

// 		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.MinCost)
// 		user := entities.User{
// 			Name:     userRequest.Name,
// 			Password: string(passwordHash),
// 			Email:    userRequest.Email,
// 		}

// 		err_regis := a.repository.Register(user)

// 		if err_regis != nil {
// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
// 		}

// 		return c.JSON(http.StatusOK, common.SuccessOperation("berhasil membuat user"))
// 	}
// }
