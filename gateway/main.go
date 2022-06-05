package main

import (
	"fmt"
	"os"

	"tiktok/base/logger"
	"tiktok/base/mymysql"
	"tiktok/base/myredis"
	"tiktok/base/snowflake"
	"tiktok/gateway/routers"
	"tiktok/setting"

	"go.uber.org/zap"
)

func initApp() error {

	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		zap.L().Fatal("load config failed, err:", zap.Error(err))
		return err
	}

	// 加载日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		zap.L().Fatal("load config failed, err:", zap.Error(err))
		fmt.Printf("init logger failed, err:%v\n", err)
		return err
	}

	// 加载 MySQL db
	if err := mymysql.InitMysql(setting.Conf.MySQLConfig); err != nil {
		zap.L().Fatal("load config failed, err:", zap.Error(err))
		return err
	}

	// 加载 Redis cache
	if err := myredis.Init(setting.Conf.RedisConfig); err != nil {
		zap.L().Fatal("init redis failed, err:", zap.Error(err))
		return err
	}
	defer myredis.Close()

	// 加载 雪花 id
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		zap.L().Fatal("init redis failed, err:", zap.Error(err))
		return err
	}

	return nil
}

func main() {

	// 检验参数是否有指定配置文件路径
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: bluebell config.yaml")
		os.Exit(1)
	}

	// 初始化各种应用
	if err := initApp(); err != nil {
		os.Exit(1)
	}

	// 启动服务
	routers.RunServer(setting.Conf.Mode)
}
