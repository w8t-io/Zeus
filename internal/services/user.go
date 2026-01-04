package services

import (
	"Zeus/internal/ctx"
	"Zeus/internal/models"
	"Zeus/internal/types"
	"Zeus/pkg/tools"
	"fmt"

	"gorm.io/gorm"
)

type (
	user struct {
		ctx *ctx.Context
	}

	UserService interface {
		Register(req interface{}) (interface{}, error)
		Login(req interface{}) (interface{}, error)
		Detail(req interface{}) (interface{}, error)
	}
)

func newUserService(ctx *ctx.Context) UserService {
	return &user{
		ctx: ctx,
	}
}

func (user *user) Register(req interface{}) (interface{}, error) {
	r := req.(*types.RequestUserRegister)
	err := r.Valid(user.ctx)
	if err != nil {
		return nil, err
	}

	var u = models.UserModel{
		Username: r.Username,
		Email:    r.Email,
		Mobile:   r.Mobile,
		Password: r.Password,
		Status:   "",
	}
	u.GenerateUserId()
	err = u.SetPassword()
	if err != nil {
		return nil, err
	}

	repo := user.ctx.Database.User()
	err = repo.Register(u)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (user *user) Login(req interface{}) (interface{}, error) {
	r := req.(*types.RequestUserLogin)
	err := r.Valid(user.ctx)
	if err != nil {
		return nil, err
	}

	repo := user.ctx.Database.User()
	data, err := repo.Detail(r.Identity)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("用户不存在")
		}

		return nil, err
	}

	if data.CheckPassword(r.Password) {
		return nil, fmt.Errorf("密码错误")
	}

	token, err := tools.GenerateToken(data.UserId, data.Username, data.Password)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (user *user) Detail(req interface{}) (interface{}, error) {
	r := req.(*types.RequestUserDetail)
	err := r.Valid()
	if err != nil {
		return nil, err
	}

	repo := user.ctx.Database.User()
	data, err := repo.Detail(r.Identity)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("用户不存在")
		}

		return nil, err
	}

	return data, nil
}
