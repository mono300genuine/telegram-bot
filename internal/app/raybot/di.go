package raybot

import (
	"github.com/google/wire"
	"raybot/internal/handler"
	"raybot/internal/service"
)

var deps = wire.NewSet(
	service.GraphSet,
	handler.GraphSet,
)

var GraphSet = wire.NewSet(
	deps,
	NewHttpServer,
	NewRayBot,
	NewApp,
)
