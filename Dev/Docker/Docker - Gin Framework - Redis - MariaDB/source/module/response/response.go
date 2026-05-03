package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
}

// SUCCESS
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(200, Response{
		StatusCode: 200,
		Success:    true,
		Message:    message,
		Data:       data,
	})
}

// SUCCESS WITH META
func Success_With_Meta(c *gin.Context, message string, data interface{}, meta interface{}) {
	c.JSON(200, Response{
		StatusCode: 200,
		Success:    true,
		Message:    message,
		Data:       data,
		Meta:       meta,
	})
}

// BAD REQUEST
func BadRequest(c *gin.Context, message string, errors interface{}) {
	c.JSON(400, Response{
		StatusCode: 400,
		Success:    false,
		Message:    message,
		Errors:     errors,
	})
}

// UNAUTHORIZED
func Unauthorized(c *gin.Context, message string) {
	c.JSON(401, Response{
		StatusCode: 401,
		Success:    false,
		Message:    message,
	})
}

// FORBIDDEN
func Forbidden(c *gin.Context, message string) {
	c.JSON(403, Response{
		StatusCode: 403,
		Success:    false,
		Message:    message,
	})
}

// NOT FOUND
func NotFound(c *gin.Context, message string) {
	c.JSON(404, Response{
		StatusCode: 404,
		Success:    false,
		Message:    message,
	})
}

// INTERNAL SERVER ERROR
func InternalError(c *gin.Context, message string, errors interface{}) {
	c.JSON(500, Response{
		StatusCode: 500,
		Success:    false,
		Message:    message,
		Errors:     errors,
	})
}

// TUỲ CHỈNH HOÀN TOÀN
func CustomError(c *gin.Context, httpStatus int, res Response) {
	res.StatusCode = httpStatus
	c.JSON(httpStatus, res)
}
