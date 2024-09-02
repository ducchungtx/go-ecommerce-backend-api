package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// success response
func SuccessResponse(c *gin.Context, code int, data interface{}) ResponseData {
	c.JSON(http.StatusOK, ResponseData{
		Code:    20001,
		Message: msg[code],
		Data:    data,
	})
	return ResponseData{} // Add a return statement
}

func ErrorResponse(c *gin.Context, code int, message string) ResponseData {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
	return ResponseData{} // Add a return statement
}
