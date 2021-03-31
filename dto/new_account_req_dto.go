package dto

import (
	"banking-app/errs"
	"strings"
)

type NewAccountRequestDto struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (dto *NewAccountRequestDto) Validate() *errs.AppError {
	if dto.Amount < 5000 {
		return errs.NewValidationError("Amount must be >= 5000")
	}

	if strings.ToLower(dto.AccountType) != "savings" && strings.ToLower(dto.AccountType) != "checking" {
		return errs.NewValidationError("Account type can be savings or checking")
	}

	return nil
}
