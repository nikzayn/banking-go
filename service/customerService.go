// Primary adapter and secondary port connection
package service

import (
	"github.com/nikzayn/banking/domain"
	"github.com/nikzayn/banking/errs"
)

// In this place, we connect primary and secondary port=
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomerById(id string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return d.repo.FindAll(status)
}

func (d DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, *errs.AppError) {
	return d.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
