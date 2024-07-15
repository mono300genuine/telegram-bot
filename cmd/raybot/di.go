//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"

	"raybot/internal/app/raybot"
)

func initRaybotApp(ctx context.Context) (raybot.App, func(), error) {
	wire.Build(raybot.GraphSet)
	return nil, nil, nil
}
