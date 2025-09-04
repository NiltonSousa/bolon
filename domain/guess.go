package domain

type Guess struct {
	ID     string  `json:"id"`
	Person *Person `json:"person" validate:"required"`
	Match  *Match  `json:"match" validate:"required"`
}

func NewGuess(id string, person *Person, match *Match) *Guess {
	return &Guess{
		ID:     id,
		Person: person,
		Match:  match,
	}
}
