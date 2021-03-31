package domain

import (
	"banking-app/dto"
	"banking-app/errs"
)

type Account struct {
	AccoundId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a *Account) ToNewAccountResponseDto() dto.NewAccountResponseDto {
	return dto.NewAccountResponseDto{AccountId: a.AccoundId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
