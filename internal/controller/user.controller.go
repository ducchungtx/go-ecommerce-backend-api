package controller

import (
	"goecommerce/internal/service"
	"goecommerce/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	// response.SuccessResponse(c, 20001, []string{"user1", "user2"})
	response.ErrorResponse(c, 20003, "Email is invalid")
}
