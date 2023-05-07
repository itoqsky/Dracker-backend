package service

import (
	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/itoqsky/money-tracker-backend/internal/storage"
)

type UserService struct {
	store      storage.User
	storeGroup storage.Group
}

func NewUserService(store storage.User, storeGroup storage.Group) *UserService {
	return &UserService{store: store, storeGroup: storeGroup}
}

func (s *UserService) GetAll(userId, groupId int) ([]core.UserInputGetAll, error) {
	// _, err := s.storeGroup.GetById(userId, groupId)
	// if err != nil {
	// 	return nil, err
	// }
	return s.store.GetAll(groupId)
}

func (s *UserService) Invite(id, groupId int, username string) error {
	_, err := s.storeGroup.GetById(id, groupId)
	if err != nil {
		return err
	}
	return s.store.Invite(groupId, username)
}

func (s *UserService) KickUser(id, groupId, kickUserId int) error {
	_, err := s.storeGroup.GetById(id, groupId)
	if err != nil {
		return err
	}

	return s.store.KickUser(groupId, kickUserId)
}
