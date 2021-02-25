// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"context"
	"factory/exam/handler"
	"factory/exam/repo"
	"factory/exam/server"
	"factory/exam/services"
)

// Injectors from buildServer.go:

func buildServer(ctx context.Context) (*server.Manager, error) {
	productRepoImp := repo.ProductRepoProvider()
	productService := services.ProductProvider(productRepoImp)
	productHandler := handler.ProductHandlerProvider(productService)
	httpServer := server.HTTPProvider(productHandler)
	manager := server.NewServerManager(httpServer)
	return manager, nil
}