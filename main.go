package main

import (
	"proj_template/pkg/config"
	"proj_template/pkg/logger"
)

func init() {
	config.InitConfig()
	logger.InitLogger()
}

func main() {
	logger.SugarLogger.Info("hello world")
}
