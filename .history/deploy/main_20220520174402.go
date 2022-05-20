package main

import (
	"fmt"
	"os"
)

// init 初始化网关层相关配置
func init() {
	// 启动
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: bluebell config.yaml")
		return
	}
}
func main() {
	
}