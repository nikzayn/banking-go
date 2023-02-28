package domain

// Define a customer
type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// Secondary port means server side ports
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
