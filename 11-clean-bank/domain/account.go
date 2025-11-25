package domain

import "errors"

var (
	ErrNotFound = errors.New("account not found")
	ErrBalance  = errors.New("insufficient balance")
)

type Account struct {
	ID      int
	Owner   string
	Balance int64 
}

type AccountRepository interface {
	GetByID(id int) (*Account, error)
	Update(account *Account) error
}