package service

import (
	"banking-app/domain"
	"banking-app/dto"
	"banking-app/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequestDto) (*domain.Account, *errs.AppError)
}

type DefaultAccountService struct {
	repository domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequestDto) (*domain.Account, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	a := domain.Account{
		AccoundId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-01 15:05:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	acc, err := s.repository.Save(a)

	if err != nil {
		return nil, err
	}

	return acc, nil
}

func NewDefaultAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repository}
}
