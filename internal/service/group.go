package service

import (
	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/itoqsky/money-tracker-backend/internal/storage"
)

type GroupService struct {
	store storage.Group
}

func NewGroupService(store storage.Group) *GroupService {
	return &GroupService{store: store}
}

func (s *GroupService) Create(userId int, group core.Group) (int, error) {
	return s.store.Create(userId, group)
}

func (s *GroupService) GetAll(userId int) ([]core.Group, error) {
	return s.store.GetAll(userId)
}

func (s *GroupService) GetById(userId, groupId int) (core.Group, error) {
	return s.store.GetById(userId, groupId)
}
