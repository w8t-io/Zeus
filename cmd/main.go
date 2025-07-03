package main

import (
	"Zeus/config"
	"Zeus/internal/cache"
	"Zeus/internal/ctx"
	"Zeus/internal/middleware"
	"Zeus/internal/repos"
	"Zeus/internal/routes"
	"Zeus/internal/services"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logc"
)

func main() {
	var do = context.Background()

	logc.Info(do, "系统初始化...")
	// 初始化配置文件
	config.InitConfig()

	// 初始化全局上下文
	ctx := ctx.NewContext(do, repos.NewEntryRepo(), cache.NewEntryCache())

	// 初始化服务逻辑层
	services.NewServices(ctx)

	// 启动服务
	mode := config.Application.Server.Mode
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	ginEngine := gin.New()
	ginEngine.Use(
		// 启用CORS中间件
		middleware.Cors(),
		// 自定义请求日志格式
		middleware.Logger(),
		gin.Recovery(),
	)
	routes.V1(ginEngine)

	logc.Info(do, fmt.Sprintf("服务启动成功 0.0.0.0:%s", config.Application.Server.Port))
	err := ginEngine.Run(":" + config.Application.Server.Port)
	if err != nil {
		logc.Error(do, "服务启动失败, err:", err)
		return
	}
}
