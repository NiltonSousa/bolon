module domain/guess

go 1.24.5

replace domain/match => ../match

replace domain/person => ../person

replace domain/team => ../team

require (
	domain/match v0.0.0-00010101000000-000000000000
	domain/person v0.0.0-00010101000000-000000000000
)

require domain/team v0.0.0-00010101000000-000000000000 // indirect
