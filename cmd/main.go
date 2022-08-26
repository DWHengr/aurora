package main

import (
	"flag"
	"github.com/DWHengr/aurora/alert"
	"github.com/DWHengr/aurora/api"
	"github.com/DWHengr/aurora/internal/service"
	config "github.com/DWHengr/aurora/pkg/config"
	"github.com/DWHengr/aurora/pkg/httpclient"
	"github.com/DWHengr/aurora/pkg/logger"
	"github.com/DWHengr/aurora/pkg/misc/email"
	"os"
	"os/signal"
	"syscall"
)

var configPath = flag.String("config", "configs/config.yml", "-config Configuration File Address")

func main() {
	flag.Parse()
	conf, err := config.NewConfig(*configPath)
	if err != nil {
		panic(err)
	}
	logger.Logger = logger.New(&conf.Log)
	if err != nil {
		panic(err)
	}
	_, err = service.NewMysqlInstanceByConn(&conf.Mysql)
	if err != nil {
		panic(err)
	}
	router, err := api.NewRouter(conf)
	if err != nil {
		panic(err)
	}

	email.NewEmail(&conf.Email)
	httpclient.NewClient(&conf.HttpClient)

	alerter := alert.NewAlerter(conf)
	alerter.Run()

	go router.Run()
	logger.Logger.Info("running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//	router.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
