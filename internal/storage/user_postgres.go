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

func (r *UserPostgres) GetAll(userId, groupId int) ([]core.UserInvitePostgres, error) {
	var users []core.UserInvitePostgres

	query := fmt.Sprintf(`SELECT u.id, u.username FROM %s u INNER JOIN %s ug ON u.id = ug.user_id WHERE ug.user_id=$1 AND ug.group_id=$2`,
		usersTable, usersGroupsTable)

	err := r.db.Select(&users, query, userId, groupId)

	return users, err
}

func (r *UserPostgres) Invite(groupId int, username string) error {
	query := fmt.Sprintf(`INSERT INTO %s (user_id, group_id) VALUES ((SELECT id FROM %s WHERE username=$1), $2)`,
		usersGroupsTable, usersTable)

	_, err := r.db.Exec(query, username, groupId)

	return err
}

func (s *UserPostgres) KickUser(id, groupId, kickUserId int) error {

}
