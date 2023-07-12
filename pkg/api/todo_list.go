package api

import "time"

type TodoListResponse struct {
	ID      uint      `json:"id"`
	ChildID uint      `json:"child_id"`
	TodoID  uint      `json:"todo_id"`
	IsDone  bool      `json:"is_done"`
	Date    time.Time `json:"date"`

	Child *ChildResponse `json:"child,omitempty"`
	ToDo  *TodoResponse  `json:"todo,omitempty"`
}

type CreateTodoListRequest struct {
	ChildID uint   `json:"child_id" binding:"required"`
	TodoID  uint   `json:"todo_id" binding:"required"`
	IsDone  bool   `json:"is_done"`
	Date    string `json:"date" binding:"required"`
}
