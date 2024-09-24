package controller

import (
	"goecommerce/internal/service"
	"goecommerce/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {

	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)
}

//
//func NewUserController() *UserController {
//
//	return &UserController{
//		userService: service.NewUserService(),
//	}
//}
//
//func (uc *UserController) GetUserByID(c *gin.Context) {
//
//	response.ErrorResponse(c, 20003, "No need!!")
//}
