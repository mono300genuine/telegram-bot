package raybot

import (
	"context"
	"fmt"
	"net"

	"github.com/getnimbus/ultrago/u_graceful"
	"github.com/getnimbus/ultrago/u_logger"
	"golang.org/x/sync/errgroup"
)

const (
	HTTP_PORT = ":8081"
)

func NewApp(
	bot *RayBot,
	httpServer *HttpServer,
) App {
	return &app{
		bot:        bot,
		httpServer: httpServer,
	}
}

type App interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type app struct {
	bot        *RayBot
	httpServer *HttpServer
}

func (a *app) Start(ctx context.Context) error {
	ctx, logger := u_logger.GetLogger(ctx)
	httpLis, err := net.Listen("tcp", HTTP_PORT)
	if err != nil {
		logger.Fatalf("failed to listen http port %s: %v", HTTP_PORT, err)
	}

	eg, childCtx := errgroup.WithContext(ctx)
	// start bot
	eg.Go(func() error {
		return u_graceful.BlockListen(childCtx, func() error {
			logger.Info("bot started!")
			return a.bot.Start()
		})
	})

	// start http server
	eg.Go(func() error {
		return u_graceful.BlockListen(childCtx, func() error {
			logger.Infof("start listening http request on %v", HTTP_PORT)
			if err := a.httpServer.Serve(httpLis); err != nil {
				return fmt.Errorf("failed to serve http: %v", err)
			}
			return nil
		})
	})

	return eg.Wait()
}

func (a *app) Stop(ctx context.Context) error {
	ctx, logger := u_logger.GetLogger(ctx)
	logger.Info("bot stopped")
	_ = a.bot.Stop()
	logger.Infof("stop listening http request on %v", HTTP_PORT)
	_ = a.httpServer.Shutdown(ctx)
	return nil
}
