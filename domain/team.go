package domain

type Team struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required"`
}

func NewTeam(id, name string) *Team {
	return &Team{
		ID:   id,
		Name: name,
	}
}
