package service

import (
	"time"

	"github.com/nikzayn/banking/domain"
	"github.com/nikzayn/banking/dto"
	"github.com/nikzayn/banking/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountRequest, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:05:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountReponseDto()
	return &response, nil
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repository}
}
