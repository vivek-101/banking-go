package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Vivek", "Delhi", "110094", "1998-03-05", "1"},
		{"1002", "Tanmay", "Delhi", "110003", "1998-05-21", "1"},
		{"1003", "Anupam", "Delhi", "110094", "1998-03-20", "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
