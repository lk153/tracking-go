package handler

import (
	"github.com/google/wire"
)

//GraphSet ...
var GraphSet = wire.NewSet(
	ProductHandlerProvider,
)
