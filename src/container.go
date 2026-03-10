package src

import (
	"github.com/MetaDandy/carpyen-service/config"
	"github.com/MetaDandy/carpyen-service/src/core/client"
	extrainformation "github.com/MetaDandy/carpyen-service/src/core/extra-information"
	"github.com/MetaDandy/carpyen-service/src/core/user"

	batchmaterial "github.com/MetaDandy/carpyen-service/src/modules/inventory/batch_material"
	batchproduct "github.com/MetaDandy/carpyen-service/src/modules/inventory/batch_product"
	batchproductmaterial "github.com/MetaDandy/carpyen-service/src/modules/inventory/batch_product_material"
	"github.com/MetaDandy/carpyen-service/src/modules/inventory/material"
	"github.com/MetaDandy/carpyen-service/src/modules/inventory/product"
	productmaterial "github.com/MetaDandy/carpyen-service/src/modules/inventory/product_material"
	"github.com/MetaDandy/carpyen-service/src/modules/inventory/purchase"
	"github.com/MetaDandy/carpyen-service/src/modules/inventory/supplier"
	"github.com/MetaDandy/carpyen-service/src/modules/inventory/warehouse"
	"github.com/MetaDandy/carpyen-service/src/modules/projects/project"
)

type Container struct {
	User             user.Handler
	Client           client.Handler
	Supplier         supplier.Handler
	Material         material.Handler
	Product          product.Handler
	BatchMaterial    batchmaterial.Handler
	BatchProduct     batchproduct.Handler
	BPM              batchproductmaterial.Handler
	PM               productmaterial.Handler
	Project          project.Handler
	ExtraInformation extrainformation.Handler
	Warehouse        warehouse.Handler
	Purchase         purchase.Handler
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

	warehouseRepo := warehouse.NewRepo(config.DB)
	warehouseService := warehouse.NewService(warehouseRepo, userRepo)
	warehouseHandler := warehouse.NewWarehouseHandler(warehouseService)

	batchMaterialRepo := batchmaterial.NewRepo(config.DB)
	batchMaterialService := batchmaterial.NewService(batchMaterialRepo, userRepo, materialRepo, warehouseRepo)
	batchMaterialHandler := batchmaterial.NewBatchMaterialHandler(batchMaterialService)

	batchProductRepo := batchproduct.NewRepo(config.DB)
	batchProductService := batchproduct.NewService(batchProductRepo, userRepo, productRepo, warehouseRepo)
	batchProductHandler := batchproduct.NewBatchProductHandler(batchProductService)

	bpmRepo := batchproductmaterial.NewRepo(config.DB)
	bpmService := batchproductmaterial.NewService(bpmRepo, userRepo, productRepo, warehouseRepo)
	bpmHandler := batchproductmaterial.NewBatchProductMaterialHandler(bpmService)

	pmRepo := productmaterial.NewRepo(config.DB)
	pmService := productmaterial.NewService(pmRepo, bpmRepo, materialRepo)
	pmHandler := productmaterial.NewHandler(pmService)

	projectRepo := project.NewRepo(config.DB)
	projectService := project.NewService(projectRepo, userRepo, clientRepo)
	projectHandler := project.NewProjectHandler(projectService)

	purchaseRepo := purchase.NewRepo(config.DB)
	purchaseService := purchase.NewService(purchaseRepo, userRepo, supplierRepo)
	purchaseHandler := purchase.NewPurchaseHandler(purchaseService)

	return &Container{
		User:             userHandler,
		Client:           clientHandler,
		Supplier:         supplierHandler,
		Material:         materialHandler,
		Product:          productHandler,
		BatchMaterial:    batchMaterialHandler,
		BatchProduct:     batchProductHandler,
		BPM:              bpmHandler,
		PM:               pmHandler,
		Project:          projectHandler,
		ExtraInformation: extrainformation.NewExtraInformationHandler(),
		Warehouse:        warehouseHandler,
		Purchase:         purchaseHandler,
	}
}
