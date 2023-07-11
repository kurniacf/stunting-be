package api

import "time"

type CreateChildRequest struct {
	Name         string `json:"name"`
	HealthStatus string `json:"health_status"`
	BirthDate    string `json:"birth_date"`
}

type ChildResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	HealthStatus string    `json:"health_status"`
	BirthDate    time.Time `json:"birth_date"`
	UserId       uint      `json:"user_id"`
}
