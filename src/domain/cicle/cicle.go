package cicle

type Cicle struct {
	id       string
	name     string
	isActive bool
}

func NewCicle(id, name string, isActive bool) *Cicle {
	return &Cicle{
		id:       id,
		name:     name,
		isActive: isActive,
	}
}
