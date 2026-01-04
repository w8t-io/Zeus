package services

import (
	"Zeus/internal/ctx"
	"context"

	"github.com/zeromicro/go-zero/core/logc"
)

var (
	User UserService
)

func NewServices(ctx *ctx.Context) {
	User = newUserService(ctx)

	logc.Info(context.Background(), "服务逻辑层初始化完成!")
}
