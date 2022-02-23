package user

import (
	"net/http"
	"sirclo/delivery/common"
	"sirclo/delivery/controllers/imageLib"
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
		// check role
		role := middlewares.GetUserRole(secret, c)
		if role != "admin" {
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
		//checking tokenId = userId
		if loginId != userId {
			return c.JSON(http.StatusUnauthorized, common.Unauthorized())
		}
		var userRequest UserRequestFormat
		// prosess binding text
		if err_bind := c.Bind(&userRequest); err_bind != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}
		theUrl := uc.repository.GetUserImageUrl(userId)
		if theUrl == "" {
			return c.JSON(http.StatusBadRequest, common.CustomResponse(500, "internal server error", "user image not found"))
		}
		// prosess binding image
		fileData, fileInfo, err_binding_image := c.Request().FormFile("image")
		if err_binding_image == nil {
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
			fileName := "user_" + strconv.Itoa(userId)

			// upload the photo
			var err_upload_photo error
			theUrl, err_upload_photo = imageLib.UploadImage("user", fileName, fileData)
			if err_upload_photo != nil {
				return c.JSON(http.StatusBadRequest, common.CustomResponse(400, "bad request", "Upload Image Failed"))
			}
		}

		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.MinCost)
		user := entities.User{
			Id:       userId,
			Name:     userRequest.Name,
			Password: string(passwordHash),
			Email:    userRequest.Email,
			ImageUrl: theUrl,
		}
		err_edit := uc.repository.EditUser(user)

		if err_edit != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}
		return c.JSON(http.StatusOK, common.SuccessOperation("berhasil mengubah data user"))
	}
}
