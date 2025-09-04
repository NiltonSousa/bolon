package domain

type Person struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Email        string `json:"email" validate:"required"`
	FavoriteTeam *Team  `json:"favorite_team" validate:"required"`
	Cicle        *Cicle `json:"cicle" validate:"required"`
	IsActive     bool   `json:"is_active" validate:"required"`
	IsAdmin      bool   `json:"is_admin" validate:"required"`
}

func NewPerson(id, firstName, lastName, email string, favoriteTeam *Team, cicle *Cicle, isActive, isAdmin bool) *Person {
	return &Person{
		ID:           id,
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		FavoriteTeam: favoriteTeam,
		Cicle:        cicle,
		IsActive:     isActive,
		IsAdmin:      isAdmin,
	}
}
