package api

import (
	"Zeus/pkg/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIError 自定义API错误类型
type APIError struct {
	Code    int
	Message string
}

// Error 实现error接口
func (e *APIError) Error() string {
	return e.Message
}

func Service(ctx *gin.Context, fu func() (interface{}, error)) {
	data, err := fu()
	if err != nil {
		if apiErr, ok := err.(*APIError); ok {
			response.Error(ctx, apiErr.Code, apiErr.Message)
		} else {
			response.Error(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}
	response.Success(ctx, data, "success")
}

func BindJson(ctx *gin.Context, req interface{}) error {
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return fmt.Errorf("[Body] 请求参数错误, %s", err.Error())
	}

	return nil
}

func BindQuery(ctx *gin.Context, req interface{}) error {
	err := ctx.ShouldBindQuery(req)
	if err != nil {
		return fmt.Errorf("[Query] 请求参数错误, %s", err.Error())
	}

	return nil
}
