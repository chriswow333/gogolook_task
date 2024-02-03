package api

import (
	reply "gogolook_task/base/apis"
	"gogolook_task/internal/task/dto"
)

type TasksReply struct {
	Tasks []*dto.TaskDto `json:"tasks"`
	Reply *reply.Reply   `json:"reply"`
}
