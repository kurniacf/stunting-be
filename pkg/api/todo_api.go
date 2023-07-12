package api

type TodoResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CreateTodoRequest struct {
	Name string `json:"name"`
}
