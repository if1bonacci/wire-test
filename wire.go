//go:build wireinject

package main

import (
	"__go/services"
	"github.com/google/wire"
)

func InitializeService() (*services.Service, error) {
	wire.Build(
		services.ProvideService,
		services.ProvideLogger,
		services.ProvideNService,
	)
	return &services.Service{}, nil
}
