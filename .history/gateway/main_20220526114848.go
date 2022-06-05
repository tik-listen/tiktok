package main

import (
	"fmt"
	"os"
	"tiktok/gateway/controller"
	"tiktok/gateway/routers"
	"tiktok/pkg/logger"
	"tiktok/pkg/mymysql"
	"tiktok/pkg/myredis"
	"tiktok/setting"

	"go.uber.org/zap"
)

func initApp() error {
	if err := setting.Init(os.Args[1]); err != nil {
		zap.L().Fatal("load config failed, err:", zap.Error(err))
		return err
	}

	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		zap.L().Fatal("load config failed, err:", zap.Error(err))
		fmt.Printf("init logger failed, err:%v\n", err)
		return err
	}

	if err := mymysql.InitMysql(setting.Conf.MySQLConfig); err != nil {
		zap.L().Fatal("load config failed, err:", zap.Error(err))
		fmt.Printf("init mysql failed, err:%v\n", err)
		return err
	}
	defer mymysql.Close()

	if err := myredis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return err
	}
	defer myredis.Close()

	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 初始化 gin 框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	// 注册路由
	r := routers.SetupRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: bluebell config.yaml")
		return
	}

}
