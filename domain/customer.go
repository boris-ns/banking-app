package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	IsActive    bool
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
