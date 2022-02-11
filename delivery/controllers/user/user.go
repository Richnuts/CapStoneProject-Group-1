package user

import (
	"net/http"
	"sirclo/delivery/common"
	"sirclo/delivery/middlewares"
	"sirclo/entities"
	userRepo "sirclo/repository/user"
	"strconv"

	echo "github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	repository userRepo.User
}

func New(user userRepo.User) *UserController {
	return &UserController{
		repository: user,
	}
}

func (uc UserController) GetProfile(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}

		result, err := uc.repository.GetUser(loginId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		return c.JSON(http.StatusOK, result)
	}
}

func (uc UserController) GetUser(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// getting the id
		userid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		result, err := uc.repository.GetUser(userid)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		return c.JSON(http.StatusOK, result)
	}
}

func (uc UserController) DeleteUser(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// getting the id
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		if userId != loginId {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		err_repo := uc.repository.DeleteUser(userId)

		if err_repo != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation("berhasil menghapus user"))
	}
}

func (uc UserController) EditUser(secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check token
		loginId := middlewares.GetUserId(secret, c)
		if loginId == 0 {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}
		// getting the id
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		if loginId != userId {
			return c.JSON(http.StatusUnauthorized, common.Unauthorized())
		}
		var userRequest UserRequestFormat

		if err_bind := c.Bind(&userRequest); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.MinCost)
		user := entities.User{
			Name:     userRequest.Name,
			Password: string(passwordHash),
			Email:    userRequest.Email,
		}

		err_edit := uc.repository.EditUser(user, userId)

		if err_edit != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation("berhasil mengubah data user"))
	}
}
