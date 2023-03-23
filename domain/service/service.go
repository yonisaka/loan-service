package service

import (
	"github.com/yonisaka/loan-service/domain/repository"
	"github.com/yonisaka/loan-service/infrastructure/persistence"
	"gorm.io/gorm"
)

// Repositories is a struct
type Repositories struct {
	Loan repository.LoanRepositoryInterface
	DB   *gorm.DB
}

// NewDBService is constructor
func NewDBService(db *gorm.DB) *Repositories {
	return &Repositories{
		Loan: persistence.NewLoanRepository(db),
		DB:   db,
	}
}
