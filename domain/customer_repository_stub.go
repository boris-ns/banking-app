package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "John Doe", "Novi Sad", "21000", "1989-01-01", true},
		{"2", "Jane Doe", "Belgrade", "11000", "1959-02-02", true},
	}

	return CustomerRepositoryStub{customers}
}
