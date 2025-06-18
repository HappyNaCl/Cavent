package service

import (
	"context"

	"github.com/HappyNaCl/Cavent/backend/internal/app/command"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/factory"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/HappyNaCl/Cavent/backend/internal/domain/repo"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/persistence/postgresdb"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/queue/tasks"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type TransactionService struct {
	transactionRepo repo.TransactionRepository

	asynq *asynq.Client
}

func NewTransactionService(db *gorm.DB, asynq *asynq.Client) *TransactionService {
	return &TransactionService{
		transactionRepo: postgresdb.NewTransactionGormRepo(db),
		asynq:           asynq,
	}
}

func (t *TransactionService) Checkout(ctx context.Context, com *command.CheckoutCommand) (*command.CheckoutCommandResult, error) {
	headerFactory := factory.NewHeaderFactory()
	detailFactory := factory.NewDetailFactory()

	header := headerFactory.CreateTransactionHeader(com.UserId, com.EventId)
	details := make([]model.TransactionDetail, len(com.Ticket))

	for i, ticket := range com.Ticket {
		detail := detailFactory.CreateTransactionDetail(header.Id, ticket.Id, ticket.Quantity)
		details[i] = detail
	}

	header.TransactionDetails = details

	err := t.transactionRepo.CreateTransaction(header)
	if err != nil {
		return nil, err
	}

	asynqTask, err := tasks.NewTicketCheckoutPayload(com.UserId, com.Ticket)
	if err != nil {
		return nil, err
	}

	if _, err := t.asynq.Enqueue(asynqTask, asynq.MaxRetry(5), asynq.Queue(tasks.TypeTicketCheckout)); err != nil {
		return nil, err
	}

	return &command.CheckoutCommandResult{}, nil
}