package service

import (
	"context"
	"errors"
	"sync"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"gogolook_task/internal/task/domain"
	"gogolook_task/internal/task/dto"
)

type TaskService interface {
	GetAllTasks(ctx context.Context) ([]*dto.TaskDto, error)
	ModifiedTask(ctx context.Context, id string, dto *dto.TaskDto) error
	CreateTask(ctx context.Context, dto *dto.TaskDto) error
	DeleteTask(ctx context.Context, id string) error
}

type impl struct {
	dig.In

	taskDomain *domain.TaskDomain
}

func New() TaskService {
	return &impl{
		taskDomain: &domain.TaskDomain{
			M:          sync.RWMutex{},
			TaskMapper: make(map[string]*domain.Task),
		},
	}
}

func (im *impl) GetAllTasks(ctx context.Context) ([]*dto.TaskDto, error) {

	tasks := im.taskDomain.GetAllTasks()

	taskDtos := []*dto.TaskDto{}
	for _, v := range tasks {
		taskDtos = append(taskDtos, toTaskDto(v))
	}
	return taskDtos, nil
}

func (im *impl) ModifiedTask(ctx context.Context, id string, taskDto *dto.TaskDto) error {

	taskDto.ID = id // trust id, not task.ID
	task, err := toTask(taskDto)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("Transfer Task failed: %s", err)
		return err
	}

	if err := im.taskDomain.ModifiedTask(id, task); err != nil {
		return err
	}

	return nil

}

func (im *impl) CreateTask(ctx context.Context, taskDto *dto.TaskDto) error {

	task, err := toTask(taskDto)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("Transfer Task failed: %s", err)
		return err
	}

	if err := im.taskDomain.CreateTask(task); err != nil {
		return err
	}

	return nil
}

func (im *impl) DeleteTask(ctx context.Context, id string) error {

	if err := im.taskDomain.DeleteTask(id); err != nil {
		logrus.WithFields(logrus.Fields{}).Info("failed to Delete")
		return err
	}
	return nil
}

func toTaskDto(task *domain.Task) *dto.TaskDto {
	return &dto.TaskDto{
		ID:     task.ID,
		Name:   task.Name,
		Status: int32(task.Status),
		Memo:   task.Memo,
	}
}

func toTask(taskDto *dto.TaskDto) (*domain.Task, error) {

	if taskDto.ID == "" {
		return nil, errors.New("Not found id")
	}

	if len(taskDto.Name) == 0 {
		return nil, errors.New("Not found name")
	}

	if check := domain.CheckTaskStatus(taskDto.Status); !check {
		return nil, errors.New("Mismatched Status")
	}

	return &domain.Task{
		ID:     taskDto.ID,
		Name:   taskDto.Name,
		Status: domain.Status(taskDto.Status),
		Memo:   taskDto.Memo,
	}, nil
}
