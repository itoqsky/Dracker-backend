package service

import (
	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/itoqsky/money-tracker-backend/internal/storage"
)

type PurchaseService struct {
	store     storage.Purchase
	userStore storage.User
}

func NewPurchaseService(store storage.Purchase, userStore storage.User) *PurchaseService {
	return &PurchaseService{store, userStore}
}

func (s *PurchaseService) Create(purchase core.Purchase) (res core.CreatePurchaseResponse, err error) {
	users, err := s.userStore.GetAll(purchase.GroupId)
	if err != nil {
		return res, err
	}
	return s.store.Create(purchase, users)
}

func (s *PurchaseService) GetAll(groupId int) ([]core.Purchase, error) {
	return s.store.GetAll(groupId)
}

func (s *PurchaseService) GetById(id int) (core.Purchase, error) {
	return s.store.GetById(id)
}

func (s *PurchaseService) Update(purchase core.Purchase) error {
	users, err := s.userStore.GetAll(purchase.GroupId)
	if err != nil {
		return err
	}
	return s.store.Update(purchase, users)
}

func (s *PurchaseService) Delete(purchase core.Purchase) error {
	users, err := s.userStore.GetAll(purchase.GroupId)
	if err != nil {
		return err
	}
	return s.store.Delete(purchase, users)
}
