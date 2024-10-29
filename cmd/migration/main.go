package main

import (
	"aphrodite-go/cmd/migration/wire"
	"aphrodite-go/pkg/config"
	"aphrodite-go/pkg/log"
	"context"
	"flag"
)

func main() {
	var envConf = flag.String("conf", "config/config.yml", "config path, eg: -conf ./config/config.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewWire(conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
