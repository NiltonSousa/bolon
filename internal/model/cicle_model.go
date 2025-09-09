package model

type CicleCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	IsActive bool   `json:"is_active" validate:"required"`
}
