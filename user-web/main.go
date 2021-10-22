package main

import (
	"fmt"
	"go.uber.org/zap"
	"shop_api/user-web/initialize"
)

func main() {
	initialize.InitLogger()
	Router := initialize.Routers()
	port := 8021
	// 使用全局的logger打印
	zap.S().Debugf("启动服务器,端口:%d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}
}
