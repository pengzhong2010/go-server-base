package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pengzhong2010/go-server-base/app/app1/internal/server/httpServer"
	"github.com/pengzhong2010/go-server-base/app/app1/internal/service"
	"github.com/pengzhong2010/go-server-base/common/pkg/client/es/es7"
	"github.com/pengzhong2010/go-server-base/common/pkg/client/mysqldb"
	"github.com/pengzhong2010/go-server-base/common/pkg/log"
)

func main() {
	// di
	_, mysqlCf, err := mysqldb.New()
	if err != nil {
		log.Error("MysqlDB err")
		return
	}
	defer mysqlCf()
	// Es7
	_, esCf, errEs := es7.New()
	if errEs != nil {
		log.Error("Es err")
		panic(errEs)
	}
	defer esCf()
	// service
	_, sCf, err := service.New()
	if err != nil {
		log.Error("Svc err")
		return
	}
	defer sCf()
	// http server
	_, hCf, err := httpServer.New()
	if err != nil {
		return
	}
	defer hCf()

	// exit
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
