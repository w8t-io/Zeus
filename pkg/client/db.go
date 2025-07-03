package client

import (
	"Zeus/config"
	"Zeus/internal/models"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	sql := config.Application.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local&timeout=%s",
		sql.Username,
		sql.Password,
		sql.Host,
		sql.Port,
		sql.Database,
		sql.Timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logc.Errorf(context.Background(), fmt.Sprintf("failed to connect database, err: %s", err.Error()))
		panic(err)
	}

	// 检查 Product 结构是否变化，变化则进行迁移
	err = db.AutoMigrate(
		&models.UserModel{},
	)
	if err != nil {
		logc.Error(context.Background(), err.Error())
		return nil
	}

	if config.Application.Server.Mode == "debug" {
		db.Debug()
	} else {
		db.Logger = logger.Default.LogMode(logger.Silent)
	}

	logc.Info(context.Background(), "Database 初始化完成!")
	return db
}
