package guess

import (
	"domain/match"
	"domain/person"
)

type Guess struct {
	id     string
	person *person.Person
	match  *match.Match
}

func NewGuess(id string, person *person.Person, match *match.Match) *Guess {
	return &Guess{
		id:     id,
		person: person,
		match:  match,
	}
}
