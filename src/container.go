package src

import (
	"github.com/MetaDandy/carpyen-service/config"
	"github.com/MetaDandy/carpyen-service/src/core/user"
	"github.com/MetaDandy/carpyen-service/src/modules/task"
)

type Container struct {
	Handler     user.Handler
	TaskHandler task.TaskHandler
}

func SetupContainer() *Container {
	userRepo := user.NewRepo(config.DB)
	userService := user.NewService(userRepo)
	userHandler := user.NewUserHandler(userService)

	taskRepo := task.NewRepo(config.DB)
	taskService := task.NewService(taskRepo)
	taskHandler := task.NewTaskHandler(taskService)

	return &Container{
		Handler:     userHandler,
		TaskHandler: taskHandler,
	}
}
