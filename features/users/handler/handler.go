package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"starter-template/features/users"
	"starter-template/helper"
	"starter-template/helper/jwt"
)

type UserHandler struct {
	service users.UserServiceInterface
	jwt     jwt.JWTInterface
}

func NewHandler(service users.UserServiceInterface, jwt jwt.JWTInterface) *UserHandler {
	return &UserHandler{
		service: service,
		jwt:     jwt,
	}
}

func (u *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterInput)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input", nil))
		}

		isValid, errors := helper.ValidateJSON(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		if !helper.ValidatePassword(input.Password) {
			errPass := []string{"Password must contain a combination letters, symbols, and numbers"}
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errPass))
		}

		var serviceInput = new(users.User)
		serviceInput.Email = input.Email
		serviceInput.PhoneNumber = input.PhoneNumber
		serviceInput.Username = input.Username
		serviceInput.Password = input.Password

		result, err := u.service.Register(*serviceInput)

		if err != nil {
			c.Logger().Info("Handler : Register Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		var response = new(RegisterResponse)
		response.Email = result.Email
		response.Username = result.Username
		response.PhoneNumber = result.PhoneNumber

		return c.JSON(http.StatusCreated, helper.FormatResponse("Register Success", response))
	}
}

func (u *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginInput)

		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input", nil))
		}

		isValid, errors := helper.ValidateJSON(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation("Invalid Format Request", errors))
		}

		result, err := u.service.Login(input.Username, input.Password)

		if err != nil {
			c.Logger().Error("Handler : Login Failed : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		var response = new(LoginResponse)
		response.Username = result.Username
		response.Token = result.Access

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Login", response))
	}
}

func (u *UserHandler) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		ext, err := u.jwt.ExtractToken(c)

		if err != nil {
			c.Logger().Error("Handler : Extract Token Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Extract Token Error", nil))
		}

		id := ext.ID

		res, err := u.service.Profile(int(id))
		if err != nil {
			c.Logger().Error("Handler : Get Profile Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Get Profile Error", nil))
		}

		var response = new(UserInfo)
		response.Email = res.Email
		response.Username = res.Username
		response.PhoneNumber = res.PhoneNumber
		response.Role = ext.Role
		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Profile", response))
	}
}
