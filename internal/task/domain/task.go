package domain

import (
	"errors"
	"sync"
)

type Status int32

const (
	InComplete Status = iota
	Complete
	EndStatus
)

type TaskDomain struct {
	M          sync.RWMutex
	TaskMapper map[string]*Task
}

type Task struct {
	ID     string
	Name   string
	Status Status
	Memo   string
}

func (d *TaskDomain) GetAllTasks() []*Task {

	d.M.RLock()
	defer d.M.RUnlock()

	tasks := []*Task{}

	for _, v := range d.TaskMapper {
		tasks = append(tasks, v)
	}

	return tasks
}

func (d *TaskDomain) ModifiedTask(id string, task *Task) error {

	d.M.Lock()
	defer d.M.Unlock()

	if _, ok := d.TaskMapper[id]; !ok {
		return errors.New("Not found Task")
	}

	if d.TaskMapper[id].Status > task.Status {
		return errors.New("The origin status is bigger than new one.")
	}

	d.TaskMapper[id] = task

	return nil
}

func (d *TaskDomain) CreateTask(task *Task) error {

	d.M.Lock()
	defer d.M.Unlock()

	if _, ok := d.TaskMapper[task.ID]; ok {
		return errors.New("Task already exist")
	}

	d.TaskMapper[task.ID] = task

	return nil
}

func (d *TaskDomain) DeleteTask(id string) error {

	d.M.Lock()
	defer d.M.Unlock()

	if _, ok := d.TaskMapper[id]; !ok {
		return errors.New("Not found Task")
	}
	delete(d.TaskMapper, id)

	return nil
}

func CheckTaskStatus(status int32) bool {
	return status < int32(EndStatus)
}
