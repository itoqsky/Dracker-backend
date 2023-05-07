package storage

import (
	"fmt"

	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (s *UserPostgres) GetAll(groupId int) ([]core.UserInputGetAll, error) {
	var users []core.UserInputGetAll

	query := fmt.Sprintf("SELECT u.id, u.username FROM %s u INNER JOIN %s ug ON u.id = ug.user_id WHERE ug.group_id=$1",
		usersTable, usersGroupsTable)

	err := s.db.Select(&users, query, groupId)

	return users, err
}

func (s *UserPostgres) Invite(groupId int, username string) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, group_id) VALUES ((SELECT id FROM %s WHERE username=$1), $2)",
		usersGroupsTable, usersTable)

	_, err := s.db.Exec(query, username, groupId)

	return err
}

func (s *UserPostgres) KickUser(groupId, kickUserId int) error {
	query := fmt.Sprintf("DELETE FROM %s ug WHERE ug.user_id=$1 AND ug.group_id=$2", usersGroupsTable)
	_, err := s.db.Exec(query, kickUserId, groupId)

	return err
}
