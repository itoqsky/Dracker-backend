package service

import (
	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/itoqsky/money-tracker-backend/internal/storage"
)

type DebtService struct {
	store storage.Debt
}

func NewDebtService(store storage.Debt) *DebtService {
	return &DebtService{store: store}
}

func (s *DebtService) GetAll(groupId int) ([]core.Debt, []core.Debt, error) {
	return s.store.GetAll(groupId)
}

func (s *DebtService) Update(debt core.Debt) error {
	return s.store.Update(debt)
}
