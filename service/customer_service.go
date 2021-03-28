package service

import "banking-app/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (dcs DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return dcs.repository.FindAll()
}

func NewDefaultCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
