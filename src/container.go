package src

import (
	"github.com/MetaDandy/carpyen-service/config"
	"github.com/MetaDandy/carpyen-service/src/core/client"
	"github.com/MetaDandy/carpyen-service/src/core/user"
	batchproductmaterial "github.com/MetaDandy/carpyen-service/src/modules/inventory/batch_product_material"
	"github.com/MetaDandy/carpyen-service/src/modules/inventory/material"
	"github.com/MetaDandy/carpyen-service/src/modules/inventory/product"
	"github.com/MetaDandy/carpyen-service/src/modules/inventory/supplier"
)

type Container struct {
	User     user.Handler
	Client   client.Handler
	Supplier supplier.Handler
	Material material.Handler
	Product  product.Handler
	BPM      batchproductmaterial.Handler
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
	materialService := material.NewService(materialRepo, userRepo)
	materialHandler := material.NewMaterialHandler(materialService)

	productRepo := product.NewRepo(config.DB)
	productService := product.NewService(productRepo, userRepo)
	productHandler := product.NewProductHandler(productService)

	bpmRepo := batchproductmaterial.NewRepo(config.DB)
	bpmService := batchproductmaterial.NewService(bpmRepo, userRepo, productRepo)
	bpmHandler := batchproductmaterial.NewBatchProductMaterialHandler(bpmService)

	return &Container{
		User:     userHandler,
		Client:   clientHandler,
		Supplier: supplierHandler,
		Material: materialHandler,
		Product:  productHandler,
		BPM:      bpmHandler,
	}
}
