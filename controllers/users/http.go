package controllers

import (
	"absency/business/users"
	"absency/controllers"
	"absency/controllers/users/request"
	"absency/controllers/users/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase users.UserUsecaseInterface
}

func NewUserController(uc users.UserUsecaseInterface) *UserController {
	return &UserController{
		usecase: uc,
	}
}

func (controller *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var userLogin request.UserLogin
	err := c.Bind(&userLogin)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	user, err := controller.usecase.Login(*userLogin.ToDomain(), ctx)
	return controllers.SuccessResponse(c, response.FromDomain(user))
}

func (controller *UserController) GetAllUsers(c echo.Context) error {
	return controllers.SuccessResponse(c, response.UserResponse{})
}
