package src

import (
	"github.com/MetaDandy/carpyen-service/config"
	"github.com/MetaDandy/carpyen-service/src/core/client"
	"github.com/MetaDandy/carpyen-service/src/core/user"
	"github.com/MetaDandy/carpyen-service/src/modules/inventory/material"
	"github.com/MetaDandy/carpyen-service/src/modules/inventory/supplier"
)

type Container struct {
	User     user.Handler
	Client   client.Handler
	Supplier supplier.Handler
	Material material.Handler
}

func SetupContainer() *Container {
	userRepo := user.NewRepo(config.DB)
	userService := user.NewService(userRepo)
	userHandler := user.NewUserHandler(userService)

	clientRepo := client.NewRepo(config.DB)
	clientService := client.NewService(clientRepo)
	clientHandler := client.NewClientHandler(clientService)

	supplierRepo := supplier.NewRepo(config.DB)
	supplierService := supplier.NewService(supplierRepo, userRepo)
	supplierHandler := supplier.NewHandler(supplierService)

	materialRepo := material.NewRepo(config.DB)
	materialService := material.NewService(materialRepo)
	materialHandler := material.NewMaterialHandler(materialService)

	return &Container{
		User:     userHandler,
		Client:   clientHandler,
		Supplier: supplierHandler,
		Material: materialHandler,
	}
}
