package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	task "gogolook_task/api/task"
	taskService "gogolook_task/internal/task/service"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(taskService.New)

	container.Provide(newServer)
	return container
}

func newServer(
	taskService taskService.TaskService,
) *gin.Engine {

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders: []string{"Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"},
	}))

	v1 := router.Group("/")
	task.NewTaskServer(v1.Group("tasks"), taskService)

	return router

}
