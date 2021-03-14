package infra

import (
	"github.com/google/wire"
)

//GraphSet ...
var GraphSet = wire.NewSet(
	InitConfiguration,
	GetConnectionPool,
)
