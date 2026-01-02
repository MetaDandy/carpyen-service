package src

import (
	"github.com/MetaDandy/carpyen-service/config"
	"github.com/MetaDandy/carpyen-service/src/core/client"
	"github.com/MetaDandy/carpyen-service/src/core/user"
)

type Container struct {
	User   user.Handler
	Client client.Handler
}

func SetupContainer() *Container {
	userRepo := user.NewRepo(config.DB)
	userService := user.NewService(userRepo)
	userHandler := user.NewUserHandler(userService)

	clientRepo := client.NewRepo(config.DB)
	clientService := client.NewService(clientRepo)
	clientHandler := client.NewClientHandler(clientService)

	return &Container{
		User:   userHandler,
		Client: clientHandler,
	}
}
