package team

type Team struct {
	id   string
	name string
}

func NewTeam(id, name string) *Team {
	return &Team{
		id:   id,
		name: name,
	}
}
