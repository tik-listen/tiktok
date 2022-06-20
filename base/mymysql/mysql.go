package mymysql

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"tiktok/source/configure"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	ConnectTimeout          = 30 * time.Second
	globalDB       *gorm.DB = nil
	err            error
)

func InitMysql(cfg *configure.MySQLConfig) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB,
	)

	globalDB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	zap.L().Info("init MySQL success:")
	return nil
}

func GetDB(ctx context.Context) *gorm.DB {
	return globalDB.WithContext(ctx)
}
