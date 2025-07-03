package api

import (
	"Zeus/internal/services"
	"Zeus/internal/types"
	"github.com/gin-gonic/gin"
)

type userController struct{}

var UserController = new(userController)

func (u *userController) API(gin *gin.RouterGroup) {
	user := gin.Group("user")
	//user.Use()
	{
		user.POST("register", u.Register)
		user.POST("login", u.Login)
		user.POST("detail", u.Detail)
	}
}

func (u *userController) Register(ctx *gin.Context) {
	r := new(types.RequestUserRegister)

	Service(ctx, func() (interface{}, error) {
		err := BindJson(ctx, r)
		if err != nil {
			return nil, err
		}
		return services.User.Register(r)
	})
}

func (u *userController) Login(ctx *gin.Context) {
	r := new(types.RequestUserLogin)

	Service(ctx, func() (interface{}, error) {
		err := BindJson(ctx, r)
		if err != nil {
			return nil, err
		}
		return services.User.Login(r)
	})
}

func (u *userController) Detail(ctx *gin.Context) {
	r := new(types.RequestUserDetail)

	Service(ctx, func() (interface{}, error) {
		err := BindJson(ctx, r)
		if err != nil {
			return nil, err
		}
		return services.User.Detail(r)
	})
}
