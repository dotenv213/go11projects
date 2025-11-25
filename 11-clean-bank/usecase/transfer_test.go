package usecase_test

import (
	"clean-bank/domain"
	"clean-bank/mocks"
	"clean-bank/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransfer_Success(t *testing.T) {
	mockRepo := new(mocks.MockAccountRepo)
	
	sender := &domain.Account{ID: 1, Owner: "Ali", Balance: 1000}
	receiver := &domain.Account{ID: 2, Owner: "Reza", Balance: 500}

	mockRepo.On("GetByID", 1).Return(sender, nil)
	mockRepo.On("GetByID", 2).Return(receiver, nil)

	mockRepo.On("Update", sender).Return(nil)
	mockRepo.On("Update", receiver).Return(nil)

	service := usecase.NewTransferService(mockRepo)

	err := service.Transfer(1, 2, 200)

	assert.NoError(t, err)
	assert.Equal(t, int64(800), sender.Balance)
	assert.Equal(t, int64(700), receiver.Balance) 
	
	mockRepo.AssertExpectations(t)
}