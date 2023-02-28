package service

import "github.com/nikzayn/banking/domain"

//In this place, we connect primary and secondary port=
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return d.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
