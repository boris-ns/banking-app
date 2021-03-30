package service

import (
	"banking-app/domain"
	"banking-app/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *errs.AppError)
	GetAllCustomersWithStatus(status string) ([]domain.Customer, *errs.AppError)
	GetCustomer(id string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (dcs DefaultCustomerService) GetAllCustomers() ([]domain.Customer, *errs.AppError) {
	return dcs.repository.FindAll()
}

func (dcs DefaultCustomerService) GetAllCustomersWithStatus(status string) ([]domain.Customer, *errs.AppError) {
	if status == "inactive" {
		status = "0"
	} else if status == "active" {
		status = "1"
	} else {
		return nil, errs.NewBadRequestError("Wrong status value")
	}

	return dcs.repository.FindAllWithStatus(status)
}

func (dcs DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return dcs.repository.FindById(id)
}

func NewDefaultCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
