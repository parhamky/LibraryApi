package Loan

import "time"

type Loan struct {
	ID         uint64    `gorm:"primaryKey"`
	UserID     uint64    `json:"user_id"`
	BookID     uint64    `json:"book_id"`
	LoanedAt   time.Time `json:"loan_date"`
	DueDate    time.Time `json:"due_date"`
	ReturnDate time.Time `json:"return_date"`
}

type Service interface {
	Add(loan *Loan) (uint64, error)
	Read(loanId *uint64) (*Loan, error)
	Update(loan *Loan, id *uint64) error
	Delete(loanId *uint64) error
}

func (Loan) TableName() string {
	return "loan"
}
