package storage

import (
	"fmt"

	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (r *GroupPostgres) Delete(users_sz, userId, groupId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	deleteUsersGroupsQuery := fmt.Sprintf("DELETE FROM %s ug WHERE ug.user_id=$1 AND ug.group_id=$2", usersGroupsTable)
	_, err = tx.Exec(deleteUsersGroupsQuery, userId, groupId)
	if err != nil {
		tx.Rollback()
		return err
	}

	if users_sz < 2 {
		deleteGroupQuery := fmt.Sprintf("DELETE FROM %s g WHERE g.id=$1", groupsTable)
		_, err = tx.Exec(deleteGroupQuery, groupId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (r *GroupPostgres) Update(userId, groupId int, input core.UpdateGroupInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		argId++
		args = append(args, *input.Name)
	}

	setQuery := ""
	if len(setValues) > 0 {
		setQuery = setValues[0]
		for i := 1; i < len(setValues); i++ {
			setQuery = fmt.Sprintf("%s, %s", setQuery, setValues[i])
		}
	}

	query := fmt.Sprintf("UPDATE %s g SET %s FROM %s ug WHERE g.id = ug.group_id AND ug.user_id=$%d AND ug.group_id=$%d",
		groupsTable, setQuery, usersGroupsTable, argId, argId+1)
	args = append(args, userId, groupId)

	logrus.Debugf("updateGroupQuery: %s", query)
	logrus.Debugf("args: %v", args)

	_, err := r.db.Exec(query, args...)

	return err
}
