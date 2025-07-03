package config

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/core/logc"
)

type Config struct {
	Server Server `json:"Server"`
	MySQL  MySQL  `json:"MySQL"`
	Redis  Redis  `json:"Redis"`
	Jwt    Jwt    `json:"Jwt"`
}

type Server struct {
	Mode string `json:"mode"`
	Port string `json:"port"`
}

type MySQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Timeout  string `json:"timeout"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

type Jwt struct {
	Expire int64 `json:"expire"`
}

var (
	Application = new(Config)
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("config/config.yaml")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		logc.Error(context.Background(), fmt.Sprintf("配置读取失败: %s", err))
		return
	}

	if err := v.Unmarshal(&Application); err != nil {
		logc.Error(context.Background(), fmt.Sprintf("配置解析失败: %s", err))
		return
	}

	logc.Info(context.Background(), "配置文件初始化完成!")
	return
}
