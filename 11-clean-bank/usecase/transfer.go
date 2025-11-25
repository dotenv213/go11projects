package usecase

import (
	"clean-bank/domain"
	"fmt"
)

type TransferService struct {
	repo domain.AccountRepository
}

func NewTransferService(repo domain.AccountRepository) *TransferService {
	return &TransferService{repo: repo}
}

func (s *TransferService) Transfer(fromID, toID int, amount int64) error {
	fromAccount, err := s.repo.GetByID(fromID)
	if err != nil {
		return err
	}

	toAccount, err := s.repo.GetByID(toID)
	if err != nil {
		return err
	}

	if fromAccount.Balance < amount {
		return domain.ErrBalance
	}

	fromAccount.Balance -= amount
	toAccount.Balance += amount

	if err := s.repo.Update(fromAccount); err != nil {
		return fmt.Errorf("failed to update sender: %w", err)
	}

	if err := s.repo.Update(toAccount); err != nil {
		return fmt.Errorf("failed to update receiver: %w", err)
	}

	return nil
}