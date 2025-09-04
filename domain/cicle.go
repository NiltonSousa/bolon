package domain

type Cicle struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" validate:"required"`
	IsActive bool   `json:"is_active" validate:"required"`
}

func NewCicle(id int64, name string, isActive bool) *Cicle {
	return &Cicle{
		ID:       id,
		Name:     name,
		IsActive: isActive,
	}
}
