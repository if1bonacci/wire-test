// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"__go/services"
)

// Injectors from wire.go:

func InitializeService() (*services.Service, error) {
	logger := services.ProvideLogger()
	nService := services.ProvideNService()
	service := services.ProvideService(logger, nService)
	return service, nil
}
