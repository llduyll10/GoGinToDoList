package controller

import (
	"GoGinToDoList/dto"
	"GoGinToDoList/service"
	"GoGinToDoList/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
	}
	userController struct {
		userService service.UserService
	}
)

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var user dto.UserCreateRequest

	if err := ctx.ShouldBind(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	result, err := c.userService.RegisterUser(ctx.Request.Context(), user)

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	ctx.JSON(http.StatusOK, res)
}
