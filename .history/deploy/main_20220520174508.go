package main

import (
	"fmt"
	"os"
)

// init 初始化网关层相关配置
func init() {
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
}
func main() {
	// 启动判断是否指定配置文件地址
	// 启动
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: tiktok config.yaml")
		return
	}
}