package config

import (
	"github.com/google/wire"
)

//GraphSet ...
var GraphSet = wire.NewSet(
	ProvideMetricPort,
)
