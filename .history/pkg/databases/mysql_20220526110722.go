package databases

import (
	"context"
	"fmt"
	"log"
	"tiktok/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

func InitMysql(cfg *setting.MySQLConfig) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB,
	)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	globalDB = db

	log.Print("db connected success")
	return nil
}

func GetDB(ctx context.Context) *gorm.DB {
	return globalDB.WithContext(ctx)
}
