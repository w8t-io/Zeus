package models

import (
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

type UserModel struct {
	gorm.Model
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Status   string `json:"status"`
}

func (user *UserModel) TableName() string {
	return "v_user"
}

func (user *UserModel) GenerateUserId() {
	limit := 8
	gid := xid.New().String()

	var xx []string
	for _, v := range gid {
		xx = append(xx, string(v))
	}

	rand.Seed(time.Now().UnixNano())
	perm := rand.Perm(len(gid))

	var id string
	for i := 0; i < limit; i++ {
		id += xx[perm[i]]
	}

	user.UserId = id
}

func (user *UserModel) SetPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)

	return nil
}

func (user *UserModel) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return err == nil
}
