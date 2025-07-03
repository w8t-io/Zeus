package api

import (
	"Zeus/pkg/response"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logc"
)

func Service(ctx *gin.Context, fu func() (interface{}, error)) {
	data, err := fu()
	if err != nil {
		logc.Error(context.Background(), err.Error())
		response.Error(ctx, err.(error).Error(), "failed")
		ctx.Abort()
		return
	} else {
		response.Success(ctx, data, "success")
	}
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
