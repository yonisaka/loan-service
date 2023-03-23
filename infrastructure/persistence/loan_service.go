package persistence

import (
	"context"

	"github.com/yonisaka/loan-service/domain/entity"
	"github.com/yonisaka/loan-service/domain/repository"

	"gorm.io/gorm"
)

type LoanRepo struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) *LoanRepo {
	return &LoanRepo{db}
}

var _ repository.LoanRepositoryInterface = &LoanRepo{}

func (u LoanRepo) Create(ctx context.Context, loan *entity.Loan) error {
	return u.db.WithContext(ctx).Create(&loan).Error
}

func (u LoanRepo) Update(ctx context.Context, id int, loan *entity.Loan) error {
	return u.db.WithContext(ctx).Model(&entity.Loan{}).Where("id = ?", id).Updates(&loan).Error
}

func (u LoanRepo) Find(ctx context.Context, id int) (*entity.Loan, error) {
	var loan entity.Loan
	err := u.db.WithContext(ctx).First(&loan, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &loan, nil
}

func (u LoanRepo) Get(ctx context.Context) ([]entity.Loan, error) {
	var loans []entity.Loan

	err := u.db.WithContext(ctx).Find(&loans).Error
	if err != nil {
		return nil, err
	}

	return loans, nil
}

func (u LoanRepo) Delete(ctx context.Context, id int) error {
	return u.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Loan{}).Error
}