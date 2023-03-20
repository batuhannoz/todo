package handler

type TodoRequest struct {
	Task string `json:"task" validate:"min=3,max=100"`
}
