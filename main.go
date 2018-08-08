package main

import (
	"short-service/utils/config"
	"fmt"
	"log"
	"syscall"
	"short-service/routers"
	"short-service/common"
	"github.com/fvbock/endless"
)

func main() {

	common.Startup()
	endless.DefaultReadTimeOut = config.ReadTimeout
	endless.DefaultWriteTimeOut = config.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	endPoint := fmt.Sprintf(":%d", config.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d\n", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server start error: %v", err)
	}
}
