package repository

import (
	"context"

	"github.com/yonisaka/loan-service/domain/entity"
)

type LoanRepositoryInterface interface {
	Create(ctx context.Context, loan *entity.Loan) error
	Get(ctx context.Context) ([]entity.Loan, error)
}
