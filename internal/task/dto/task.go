package dto

type TaskDto struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Status int32  `json:"status"`
	Memo   string `json:"memo,omitempty"`
}
