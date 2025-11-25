package mocks

import (
	"clean-bank/domain"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepo struct {
	mock.Mock
}

func (m *MockAccountRepo) GetByID(id int) (*domain.Account, error) {
	args := m.Called(id) 
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Account), args.Error(1)
}

func (m *MockAccountRepo) Update(account *domain.Account) error {
	args := m.Called(account)
	return args.Error(0)
}