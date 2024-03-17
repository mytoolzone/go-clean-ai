package resp

import (
	"net/http"

	"go-clean/internal/app_err"
	"go-clean/internal/entity/dto"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code       string          `json:"code"`
	Data       interface{}     `json:"data"`
	Doc        string          `json:"doc,omitempty"`
	Message    string          `json:"message,omitempty"`
	Pagination *dto.Pagination `json:"pagination,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Data: data,
	})
}

// SuccessWithPagination 请求成功需要分页
func SuccessWithPagination(c *gin.Context, data interface{}, pagination *dto.Pagination) {
	c.JSON(http.StatusOK, Response{
		Data:       data,
		Pagination: pagination,
	})
}

// Fail 请求失败时候调用
func Fail(c *gin.Context, err error) {
	if err == nil {
		return
	}

	appErr, ok := err.(*app_err.AppError)
	if !ok {
		c.JSON(http.StatusInternalServerError, Response{
			Message: err.Error(),
		})
		return
	}

	var httpCode int
	var doc string
	switch appErr.Code {
	case app_err.ErrBadRequest, app_err.ErrConflict:
		httpCode = http.StatusBadRequest
	case app_err.ErrExpiredToken, app_err.ErrInvalidToken, app_err.ErrInvalidCredentials, app_err.ErrPermissionDenied:
		httpCode = http.StatusUnauthorized
	case app_err.ErrNotFound:
		httpCode = http.StatusNotFound
	case app_err.ErrFileReadFailed:
		httpCode = http.StatusInternalServerError
	default:
		httpCode = http.StatusInternalServerError
	}

	c.JSON(httpCode, Response{
		Code:    appErr.Code,
		Doc:     doc,
		Message: appErr.Message,
	})
	return
}
