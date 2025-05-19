package main

import (
	log "github.com/sirupsen/logrus"
	"order-api/config"
	"order-api/di"
	"order-api/server"
	"order-api/utils"
)

func main() {
	server.Run()
}

func init() {
	initLog()
}

func initLog() {
	cfg := di.Get[*config.Config]("config")
	logLvl, err := log.ParseLevel(cfg.LoggerConfig.Level)

	utils.IsError(err, config.ErrParseLog)
	log.SetLevel(logLvl)
}
