package service

import (
	"context"
	"gogolook_task/internal/task/dto"
	"testing"
)

func TestGetAllTasks(t *testing.T) {
	svc := New()
	ctx := context.Background()

	dtos, err := svc.GetAllTasks(ctx)

	if err != nil || len(dtos) != 0 {
		t.Error("Failed to test")
	} else {
		t.Log("Pass")
	}

	err = svc.CreateTask(ctx, &dto.TaskDto{
		ID:     "id1",
		Name:   "name",
		Status: 0,
		Memo:   "meme",
	})

	if err != nil {
		t.Error("Failed to test")
	} else {
		t.Log("Pass")
	}

	dtos, err = svc.GetAllTasks(ctx)
	if len(dtos) != 1 && err != nil {
		t.Error("Failed to test")
	} else {
		t.Log("Pass")
	}

}

func TestModifiedTask(t *testing.T) {
	svc := New()
	ctx := context.Background()

	mockDto := &dto.TaskDto{
		ID:     "id1",
		Name:   "name",
		Status: 0,
		Memo:   "meme",
	}

	err := svc.ModifiedTask(ctx, "id1", mockDto)
	if err == nil {
		t.Error("Failed to test")
	} else {
		t.Log("Pass")
	}

	err = svc.CreateTask(ctx, mockDto)
	if err == nil {
		t.Log("Pass")
	} else {
		t.Error("Failed to test")
	}

	err = svc.ModifiedTask(ctx, "id1", mockDto)
	if err == nil {
		t.Log("Pass")
	} else {
		t.Error("Failed to test")
	}

}

func TestCreateTask(t *testing.T) {
	svc := New()
	ctx := context.Background()

	err := svc.CreateTask(ctx, &dto.TaskDto{
		ID:     "id",
		Name:   "name",
		Status: 0,
		Memo:   "meme",
	})

	if err == nil {
		t.Log("Pass")
	} else {
		t.Error("Failed to test")
	}

	err = svc.CreateTask(ctx, &dto.TaskDto{
		ID:     "id",
		Name:   "name",
		Status: 0,
		Memo:   "meme",
	})

	if err == nil {
		t.Error("Failed to test")
	} else {
		t.Log("Pass")
	}

	err = svc.CreateTask(ctx, &dto.TaskDto{
		ID: "without name",
	})

	if err == nil {
		t.Error("Failed to test")
	} else {
		t.Log("Pass ")
	}

	err = svc.CreateTask(ctx, &dto.TaskDto{
		Name: "without id",
	})

	if err == nil {
		t.Error("Failed to test")
	} else {
		t.Log("Pass ")
	}

}

func TestDeleteTask(t *testing.T) {

	svc := New()
	ctx := context.Background()

	err := svc.DeleteTask(ctx, "id")

	if err == nil {
		t.Error("Failed to test")
	} else {
		t.Log("Pass")
	}

	err = svc.CreateTask(ctx, &dto.TaskDto{
		ID:     "id",
		Name:   "name",
		Status: 0,
		Memo:   "meme",
	})

	if err == nil {
		t.Log("Pass")
	} else {
		t.Error("Failed to test")
	}

	err = svc.DeleteTask(ctx, "id")

	if err == nil {
		t.Log("Pass")
	} else {
		t.Error("Failed to test")
	}
}
