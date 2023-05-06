package storage

import (
	"fmt"

	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/jmoiron/sqlx"
)

type GroupPostgres struct {
	db *sqlx.DB
}

func NewGroupPostgres(db *sqlx.DB) *GroupPostgres {
	return &GroupPostgres{db: db}
}

func (r *GroupPostgres) Create(userId int, group core.Group) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createGroupQuery := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", groupsTable)
	row := tx.QueryRow(createGroupQuery, group.Name)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersGroupsQuery := fmt.Sprintf("INSERT INTO %s (user_id, group_id) values ($1, $2)", usersGroupsTable)
	_, err = tx.Exec(createUsersGroupsQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *GroupPostgres) GetAll(userId int) ([]core.Group, error) {
	var groups []core.Group

	query := fmt.Sprintf("SELECT g.id, g.name FROM %s g INNER JOIN %s ug ON g.id = ug.group_id WHERE ug.user_id=$1",
		groupsTable, usersGroupsTable)
	err := r.db.Select(&groups, query, userId)

	return groups, err
}

func (r *GroupPostgres) GetById(userId, groupId int) (core.Group, error) {
	var group core.Group

	query := fmt.Sprintf("SELECT g.id, g.name FROM %s g INNER JOIN %s ug ON g.id = ug.group_id WHERE ug.user_id=$1 AND ug.group_id=$2",
		groupsTable, usersGroupsTable)
	err := r.db.Get(&group, query, userId, groupId)

	return group, err
}
