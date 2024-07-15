package main

import (
	"context"
	"os"
	"time"

	"github.com/getnimbus/ultrago/u_graceful"
	"github.com/getnimbus/ultrago/u_logger"
	"github.com/sirupsen/logrus"
	"raybot/internal/conf"
)

func init() {
	os.Setenv("TZ", "UTC")
	_, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	if conf.Config.IsDebug() {
		u_logger.WithFormatter(logrus.DebugLevel)
	} else {
		u_logger.WithFormatter(logrus.InfoLevel)
	}
}

func main() {
	ctx, logger := u_logger.GetLogger(u_graceful.NewCtx())

	// load config env
	if err := conf.LoadConfig("."); err != nil {
		logger.Fatalf("cannot load config: %v", err)
	}
	logger.Info("env is: ", conf.Config.Env)

	// init app
	app, cleanup, err := initRaybotApp(ctx)
	if err != nil {
		logger.Fatalf("failed to init app: %v", err)
	}
	defer func() {
		shutDownErr := app.Stop(context.Background())
		logger.Infof("service is shutdown with err: %v", shutDownErr)
		cleanup() // close connection such as db, redis, kafka,...
	}()

	err = app.Start(ctx)
	if err != nil {
		logger.Fatalf("failed to start app: %v", err)
	}
}
