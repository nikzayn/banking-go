package domain

//Secondary adapter with integration of port
type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id:          "1001",
			Name:        "Nikhil",
			City:        "Delhi",
			Zipcode:     "110063",
			DateOfBirth: "1997-02-02",
			Status:      "1",
		},
		{
			Id:          "1002",
			Name:        "Pooja",
			City:        "Delhi",
			Zipcode:     "110083",
			DateOfBirth: "1999-02-19",
			Status:      "1",
		},
	}

	return CustomerRepositoryStub{customers}
}
