package apis

type Status int32

const (
	Ok Status = iota
	Fail
)

type Reply struct {
	Status Status `json:"status"`
	Error  string `json:"error,omitempty"`
}
