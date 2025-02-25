package Loan

import "errors"

type MockService struct {
	Loans map[uint64]*Loan
}

func (m *MockService) Add(loan *Loan) (uint64, error) {
	newID := uint64(len(m.Loans) + 1)
	loan.ID = newID
	m.Loans[newID] = loan
	return newID, nil
}

func (m *MockService) Read(id *uint64) (*Loan, error) {
	if loan, ok := m.Loans[*id]; ok {
		return loan, nil
	}
	return nil, errors.New("loan not found")
}

func (m *MockService) Update(loan *Loan, id *uint64) error {
	if _, ok := m.Loans[*id]; ok {
		loan.ID = *id
		m.Loans[*id] = loan
		return nil
	} else {
		return errors.New("loan not found")
	}
}

func (m *MockService) Delete(loanId *uint64) error {
	if _, ok := m.Loans[*loanId]; ok {
		delete(m.Loans, *loanId)
		return nil
	} else {
		return errors.New("loan not found")
	}
}

func GetMockService() *MockService {
	return &MockService{Loans: make(map[uint64]*Loan)}
}
