package domain

import (
	"banking-app/dto"
	"banking-app/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) ToDto() dto.CustomerDto {
	return dto.CustomerDto{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusToText(),
	}
}

func (c Customer) statusToText() string {
	if c.Status == "1" {
		return "active"
	}

	return "inactive"
}

func CustomersToDto(customers []Customer) []dto.CustomerDto {
	dtos := make([]dto.CustomerDto, 0, cap(customers))

	for i := range customers {
		c := customers[i].ToDto()
		dtos = append(dtos, c)
	}

	return dtos
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindAllWithStatus(status string) ([]Customer, *errs.AppError)
	FindById(id string) (*Customer, *errs.AppError)
}
