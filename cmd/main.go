package main

import (
	"aurora/api"
	"flag"
	"os"
	"os/signal"
	config "aurora/internal/config"
	"aurora/internal/logger"
	"syscall"
)

var configPath = flag.String("config", "../configs/config.yml", "-config Configuration File Address")

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
