// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"user-service/internal/biz"
	"user-service/internal/conf"
	"user-service/internal/data"
	"user-service/internal/server"
	"user-service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserDataRepo(dataData, logger)
	userUsecase := biz.NewUserUseCase(userRepo, logger)
	userService := service.NewuserService(userUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, userService, logger)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	r := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, r)
	return app, func() {
		cleanup()
	}, nil
}
