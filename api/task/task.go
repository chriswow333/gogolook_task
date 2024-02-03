package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"gogolook_task/base/apis"
	taskService "gogolook_task/internal/task/service"

	taskDto "gogolook_task/internal/task/dto"
)

type server struct {
	dig.In

	taskService taskService.TaskService
}

func NewTaskServer(
	rg *gin.RouterGroup,

	taskService taskService.TaskService,
) {
	s := &server{
		taskService: taskService,
	}

	rg.GET("/", s.getAllTasks)
	rg.POST("/", s.createTask)
	rg.PUT("/:id", s.modifiedTask)
	rg.DELETE("/:id", s.deleteTask)

}

func (s *server) getAllTasks(ctx *gin.Context) {

	logrus.WithFields(logrus.Fields{}).Info("Request")

	tasks, err := s.taskService.GetAllTasks(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("taskService.GetAllTasks failed: ", err)

		ctx.JSON(http.StatusInternalServerError,
			&TasksReply{
				Reply: &apis.Reply{
					Status: apis.Fail,
					Error:  "GetAllTasks failed",
				},
			},
		)
		return
	}

	ctx.JSON(http.StatusOK,
		&TasksReply{
			Reply: &apis.Reply{
				Status: apis.Ok,
			},
			Tasks: tasks,
		},
	)
}

func (s *server) createTask(ctx *gin.Context) {
	logrus.WithFields(logrus.Fields{}).Info("Request")

	var taskDto taskDto.TaskDto
	ctx.BindJSON(&taskDto)

	if err := s.taskService.CreateTask(ctx, &taskDto); err != nil {
		logrus.WithFields(logrus.Fields{}).Error("taskService.CreateTask failed: ", err)

		ctx.JSON(http.StatusInternalServerError,
			&TasksReply{
				Reply: &apis.Reply{
					Status: apis.Fail,
					Error:  "CreateTask failed",
				},
			},
		)
		return
	}

	ctx.JSON(http.StatusOK,
		&TasksReply{
			Reply: &apis.Reply{
				Status: apis.Ok,
			},
		},
	)
}

func (s *server) deleteTask(ctx *gin.Context) {

	logrus.WithFields(logrus.Fields{}).Info("Request")

	id := ctx.Param("id")

	if err := s.taskService.DeleteTask(ctx, id); err != nil {
		logrus.WithFields(logrus.Fields{}).Error("taskService.DeleteTask failed: ", err)

		ctx.JSON(http.StatusInternalServerError,
			&TasksReply{
				Reply: &apis.Reply{
					Status: apis.Fail,
					Error:  "DeleteTask failed",
				},
			},
		)
		return
	}

	ctx.JSON(http.StatusOK,
		&TasksReply{
			Reply: &apis.Reply{
				Status: apis.Ok,
			},
		},
	)
}

func (s *server) modifiedTask(ctx *gin.Context) {
	logrus.WithFields(logrus.Fields{}).Info("Request")

	id := ctx.Param("id")

	var taskDto taskDto.TaskDto
	ctx.BindJSON(&taskDto)

	if err := s.taskService.ModifiedTask(ctx, id, &taskDto); err != nil {
		logrus.WithFields(logrus.Fields{}).Error("taskService.ModifiedTask failed: ", err)

		ctx.JSON(http.StatusInternalServerError,
			&TasksReply{
				Reply: &apis.Reply{
					Status: apis.Fail,
					Error:  "ModifiedTask failed",
				},
			},
		)
		return
	}

	ctx.JSON(http.StatusOK,
		&TasksReply{
			Reply: &apis.Reply{
				Status: apis.Ok,
			},
		},
	)
}
