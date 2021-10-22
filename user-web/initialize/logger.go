package initialize

import "go.uber.org/zap"

func InitLogger() {
	// 设置全局的logger
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
