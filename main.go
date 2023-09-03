package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/sabahtalateh/gic"
	"github.com/sabahtalateh/gicex/internal/config"
	"github.com/sabahtalateh/gicex/internal/services"
)

func main() {
	err := gic.ConfigureGlobalContainer(
		gic.WithDump(gic.WithDumpDir("./dump")),
	)
	check(err)

	err = gic.Init()
	check(err)

	conf, err := gic.GetE[config.Config]()
	check(err)

	start(conf.App)

	s := gic.Get[*services.SomeService]()
	some, err := s.GetSome()
	if err != nil {
		log.Printf("err: %s", err)
	}
	println(some)

	// Wait Ctrl+C
	println("Ctrl+C to stop")
	intC := make(chan os.Signal, 1)
	signal.Notify(intC, os.Interrupt)
	<-intC

	stop(conf.App.StopTimeout)

	println("stopped")
}

func start(appConf config.App) {
	startCtx, startCnl := context.WithTimeout(context.Background(), appConf.StartTimeout)
	defer startCnl()
	err := gic.Start(startCtx)
	if err != nil {
		println("not started. stopping")
		stop(appConf.StopTimeout)
		panic(err)
	}
}

func stop(to time.Duration) {
	stopCtx, stopCnl := context.WithTimeout(context.Background(), to)
	defer stopCnl()
	err := gic.Stop(stopCtx)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
