package main

import (
	"context"
	"flag"
	"fmt"

	"aphrodite-go/cmd/server/wire"
	"aphrodite-go/pkg/config"
	"aphrodite-go/pkg/log"

	"go.uber.org/zap"
)

// @title           Aphrodite API
// @description     API Description
// @version         1.0.0

// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
func main() {
	var envConf = flag.String("conf", "config.yml", "config path, eg: -conf ./config.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewWire(conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	logger.Info("===============================")
	logger.Info("Listening on", zap.String("host", fmt.Sprintf("http://%s:%d", conf.GetString("http.host"), conf.GetInt("http.port"))))
	logger.Info("Docs addr", zap.String("addr", fmt.Sprintf("http://%s:%d/swagger-ui/index.html", conf.GetString("http.host"), conf.GetInt("http.port"))))
	logger.Info("===============================")
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
