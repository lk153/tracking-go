package server

import (
	"github.com/google/wire"

	"factory/exam/config"
	"factory/exam/handler"
	"factory/exam/repo"
	"factory/exam/services"
)

//ServerDeps ...
var ServerDeps = wire.NewSet(
	config.GraphSet,
	handler.GraphSet,
	services.GraphSet,
	repo.GraphSet,
)

//GraphSet ...
var GraphSet = wire.NewSet(
	ServerDeps,
	HTTPProvider,
	NewMetricServer,
	NewServerManager,
)
