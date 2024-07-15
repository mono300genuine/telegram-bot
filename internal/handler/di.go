package handler

import (
	"github.com/google/wire"
)

var GraphSet = wire.NewSet(
	NewPoolHandler,
)
