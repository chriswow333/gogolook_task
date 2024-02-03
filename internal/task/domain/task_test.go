package domain

import (
	"sync"
	"testing"
)

func TestDeleteTask(t *testing.T) {

	testDomain := &TaskDomain{
		M:          sync.RWMutex{},
		TaskMapper: make(map[string]*Task),
	}

	actualErr := testDomain.DeleteTask("id")

	if actualErr == nil {
		t.Errorf("Failed to test %s", actualErr)
	} else {
		t.Log("Pass")
	}

	testDomain.TaskMapper["id"] = &Task{
		ID:   "id",
		Name: "name",
		Memo: "meme",
	}

	actualErr = testDomain.DeleteTask("id")

	if actualErr == nil {
		t.Log("Pass")
	} else {
		t.Errorf("Failed to test %s", actualErr)
	}

}

func TestCreateTask(t *testing.T) {

	testDomain := &TaskDomain{
		M:          sync.RWMutex{},
		TaskMapper: make(map[string]*Task),
	}

	actualErr := testDomain.CreateTask(&Task{
		ID:   "id",
		Name: "name",
		Memo: "meme",
	})

	if actualErr == nil {
		t.Log("Pass")
	} else {
		t.Errorf("Failed to test %s", actualErr)
	}

	actualErr = testDomain.CreateTask(&Task{
		ID:   "id",
		Name: "name",
		Memo: "meme",
	})

	if actualErr != nil {
		t.Logf("Pass %s", actualErr)
	} else {
		t.Errorf("Failed to test %s", actualErr)
	}

}

func TestModifiedTask(t *testing.T) {
	testDomain := &TaskDomain{
		M:          sync.RWMutex{},
		TaskMapper: make(map[string]*Task),
	}

	actualErr := testDomain.ModifiedTask("id", &Task{
		ID:   "id",
		Name: "name",
		Memo: "meme",
	})

	if actualErr != nil {
		t.Logf("Pass %s", actualErr)
	} else {
		t.Error("Failed to test")
	}

	testDomain.TaskMapper["id"] = &Task{
		ID:   "id",
		Name: "name",
		Memo: "meme",
	}

	actualErr = testDomain.ModifiedTask("id", &Task{
		ID:   "id",
		Name: "name",
		Memo: "momo",
	})

	if actualErr == nil {
		t.Log("Pass")
	} else {
		t.Errorf("Failed to test %s", actualErr)
	}

}

func TestGetAllTasks(t *testing.T) {
	testDomain := &TaskDomain{
		M:          sync.RWMutex{},
		TaskMapper: make(map[string]*Task),
	}

	actualTasks := testDomain.GetAllTasks()
	if len(actualTasks) != 0 {
		t.Error("Failed to test")
	} else {
		t.Log("Pass")
	}

	testDomain.TaskMapper["ok"] = &Task{
		ID:   "ok",
		Name: "name",
		Memo: "meme",
	}

	actualTasks = testDomain.GetAllTasks()
	if len(actualTasks) != 1 {
		t.Error("Failed to test")
	} else {
		t.Log("Pass")
	}

}

func TestCheckTaskStatus(t *testing.T) {

	check := CheckTaskStatus(int32(Complete))
	if check {
		t.Log("Pass")
	} else {
		t.Error("Failed to test")

	}

	check = CheckTaskStatus(int32(EndStatus))
	if !check {
		t.Log("Pass")
	} else {
		t.Error("Failed to test")
	}

}
