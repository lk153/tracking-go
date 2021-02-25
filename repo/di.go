package repo

import (
	"github.com/google/wire"
)

//GraphSet ...
var GraphSet = wire.NewSet(
	ProductRepoProvider,
	wire.Bind(new(ProductRepoInterface), new(*ProductRepoImp)),
)
