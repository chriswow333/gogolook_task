package routes

import (
	"encoding/json"

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

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(jsonLoggerMiddleware())
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

func jsonLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			log := make(map[string]interface{})

			log["status_code"] = params.StatusCode
			log["path"] = params.Path
			log["method"] = params.Method
			log["start_time"] = params.TimeStamp.Format("2006/01/02 15:04:05")
			log["remote_addr"] = params.ClientIP
			log["response_time"] = params.Latency.String()

			s, _ := json.Marshal(log)
			return string(s) + "\n"
		},
	)
}
