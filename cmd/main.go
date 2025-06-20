package main

import (
	log "github.com/sirupsen/logrus"
	"order-api/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go server.Run()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("quit")
}

func gracefulShutdown() {

}
