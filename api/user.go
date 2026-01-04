package api

import (
	"Zeus/internal/services"
	"Zeus/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userAPI struct{}

var User = new(userAPI)

func (u *userAPI) API(gin *gin.RouterGroup) {
	user := gin.Group("user")
	//user.Use()
	{
		user.POST("register", u.Register)
		user.POST("login", u.Login)
		user.POST("detail", u.Detail)
	}
}

func (u *userAPI) Register(ctx *gin.Context) {
	r := new(types.RequestUserRegister)

	Service(ctx, func() (interface{}, error) {
		err := BindJson(ctx, r)
		if err != nil {
			return nil, &APIError{Code: http.StatusBadRequest, Message: "Invalid request body"}
		}
		return services.User.Register(r)
	})
}

func (u *userAPI) Login(ctx *gin.Context) {
	r := new(types.RequestUserLogin)

	Service(ctx, func() (interface{}, error) {
		err := BindJson(ctx, r)
		if err != nil {
			return nil, &APIError{Code: http.StatusBadRequest, Message: "Invalid request body"}
		}
		return services.User.Login(r)
	})
}

func (u *userAPI) Detail(ctx *gin.Context) {
	r := new(types.RequestUserDetail)

	Service(ctx, func() (interface{}, error) {
		err := BindJson(ctx, r)
		if err != nil {
			return nil, &APIError{Code: http.StatusBadRequest, Message: "Invalid request body"}
		}
		return services.User.Detail(r)
	})
}
