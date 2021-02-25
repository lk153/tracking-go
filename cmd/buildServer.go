// +build wireinject

package main

import (
	"context"

	"factory/exam/server"

	"github.com/google/wire"
)

func buildServer(ctx context.Context) (*server.Manager, error) {
	panic(wire.Build(server.GraphSet))
}
