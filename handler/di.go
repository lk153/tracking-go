package handler

import (
	"github.com/google/wire"

	services_pb "github.com/lk153/proto-tracking-gen/go/services"
)

//GraphSet ...
var GraphSet = wire.NewSet(
	NewProductPBHandler,
	wire.Bind(new(services_pb.ProductServiceServer), new(*ProductPBHandler)),

	NewTaskPBHandler,
	wire.Bind(new(services_pb.TaskServiceServer), new(*TaskPBHandler)),
)
