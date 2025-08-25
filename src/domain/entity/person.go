package person

import (
	"domain/cicle"
	"domain/team"
)

type Person struct {
	id           string
	firstName    string
	lastName     string
	email        string
	favoriteTeam *team.Team
	cicle        *cicle.Cicle
	isActive     bool
	isAdmin      bool
}

func NewPerson(id, firstName, lastName, email string, favoriteTeam *team.Team, cicle *cicle.Cicle, isActive, isAdmin bool) *Person {
	return &Person{
		id:           id,
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		favoriteTeam: favoriteTeam,
		cicle:        cicle,
		isActive:     isActive,
		isAdmin:      isAdmin,
	}
}
