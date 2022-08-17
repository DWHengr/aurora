package main

import (
	"flag"
	"github.com/DWHengr/aurora/api"
	"github.com/DWHengr/aurora/internal/alertcore"
	config "github.com/DWHengr/aurora/pkg/config"
	"github.com/DWHengr/aurora/pkg/logger"
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
	router, err := api.NewRouter(conf)
	if err != nil {
		panic(err)
	}

	alertcore.Run(&conf.Alert)

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
