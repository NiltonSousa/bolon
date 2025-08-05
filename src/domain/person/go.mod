module domain/person

go 1.24.5

replace domain/team => ../team

require (
	domain/cicle v0.0.0-00010101000000-000000000000
	domain/team v0.0.0-00010101000000-000000000000
)

replace domain/cicle => ../cicle
