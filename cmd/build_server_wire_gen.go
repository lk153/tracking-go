// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire gen -tags "dynamic"
//+build !wireinject

package main

import (
	"context"
	"factory/exam/config"
	"factory/exam/handler"
	"factory/exam/infra"
	"factory/exam/repo/cache"
	"factory/exam/repo/mysql"
	"factory/exam/server"
	"factory/exam/services"
)

// Injectors from buildServer.go:

func buildServer(ctx context.Context) (*server.Manager, error) {
	configuration := infra.InitConfiguration()
	connPool, err := infra.GetConnectionPool(configuration)
	if err != nil {
		return nil, err
	}
	productMySQLRepo := mysql.NewProductMySQLRepo(connPool)
	redisCache := cache.NewRedisCacheRepo()
	productService := services.ProductProvider(productMySQLRepo, redisCache)
	productPBHandler := handler.NewProductPBHandler(productService)
	taskMySQLRepo := mysql.NewTaskMySQLRepo(connPool)
	taskService := services.TaskProvider(taskMySQLRepo, redisCache)
	taskPBHandler := handler.NewTaskPBHandler(taskService)
	httpServer, err := server.HTTPProvider(ctx, productPBHandler, taskPBHandler)
	if err != nil {
		return nil, err
	}
	metricPort := config.ProvideMetricPort()
	metricServer, err := server.NewMetricServer(metricPort)
	if err != nil {
		return nil, err
	}
	kafkaConsumer, err := server.NewKafkaConsumer(productService, taskService, productMySQLRepo)
	if err != nil {
		return nil, err
	}
	manager := server.NewServerManager(httpServer, metricServer, kafkaConsumer)
	return manager, nil
}
