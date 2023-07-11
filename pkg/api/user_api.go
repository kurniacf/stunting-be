package api

type GetUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	ID    uint
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
