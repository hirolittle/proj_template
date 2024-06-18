package main

import (
	"proj_template/config"
	"proj_template/logger"
)

func init() {
	config.InitConfig()
	logger.InitLogger()
}

func main() {
	logger.SugarLogger.Info("hello world")
}
