package storage

import (
	"fmt"

	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/jmoiron/sqlx"
)

type PurchasePostgres struct {
	db *sqlx.DB
}

func NewPurchasePostgres(db *sqlx.DB) *PurchasePostgres {
	return &PurchasePostgres{db: db}
}

func (r *PurchasePostgres) Create(purchase core.Purchase, users []core.UserInputGetAll) (res core.CreatePurchaseResponse, err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return res, err
	}

	insertInto := fmt.Sprintf("group_id, amount, buyer_id, description")
	insertValues := fmt.Sprint("$1, $2, $3, $4")

	args := make([]interface{}, 0)
	args = append(args, purchase.GroupId, purchase.Amount, purchase.BuyerId, purchase.Description)

	// var id int
	insertPurchaseQuery := fmt.Sprintf("INSERT INTO %s (%s) values (%s) RETURNING id, timestamp",
		purchasesTable, insertInto, insertValues)
	row := tx.QueryRow(insertPurchaseQuery, args...)
	if err := row.Scan(&res.ID, &res.Timestamp); err != nil {
		tx.Rollback()
		return res, err
	}

	for _, user := range users {
		if user.Id == purchase.BuyerId {
			continue
		}
		args = make([]interface{}, 0)
		args = append(args, purchase.BuyerId, user.Id, purchase.Amount/float32(len(users)))
		updateDebtQuery := fmt.Sprintf(`INSERT INTO %s (creditor_id, debtor_id, amount) VALUES ($1, $2, $3)
			ON CONFLICT (creditor_id, debtor_id) DO UPDATE SET amount = %s.amount + $3`, debtsTable, debtsTable)
		_, err = tx.Exec(updateDebtQuery, args...)
		if err != nil {
			tx.Rollback()
			return res, err
		}
	}
	return res, tx.Commit()
}

func (r *PurchasePostgres) GetAll(groupId int) ([]core.Purchase, error) {
	var purchases []core.Purchase

	query := fmt.Sprintf("SELECT id, group_id, amount, buyer_id, description, timestamp FROM %s WHERE group_id=$1", purchasesTable)
	err := r.db.Select(&purchases, query, groupId)

	return purchases, err
}

func (r *PurchasePostgres) GetById(id int) (core.Purchase, error) {
	var purchase core.Purchase
	query := fmt.Sprintf("SELECT id, group_id, amount, buyer_id, description, timestamp FROM %s WHERE id=$1", purchasesTable)
	err := r.db.Get(&purchase, query, id)

	return purchase, err
}

func (r *PurchasePostgres) Update(purchase core.Purchase, users []core.UserInputGetAll) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	newAmount := purchase.Amount
	query := fmt.Sprintf("SELECT amount FROM %s WHERE id=$1", purchasesTable)
	row := tx.QueryRow(query, purchase.ID)
	if err := row.Scan(&purchase.Amount); err != nil {
		tx.Rollback()
		return err
	}

	purchase.Amount = newAmount - purchase.Amount

	setValues := fmt.Sprintf("amount=$1, description=$2")
	query = fmt.Sprintf("UPDATE %s SET %s WHERE id=$3", purchasesTable, setValues)
	_, err = tx.Exec(query, purchase.Amount, purchase.Description, purchase.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, user := range users {
		if user.Id == purchase.BuyerId {
			continue
		}
		args := make([]interface{}, 0)
		args = append(args, purchase.BuyerId, user.Id, purchase.Amount/float32(len(users)))
		updateDebtQuery := fmt.Sprintf(`UPDATE %s SET amount = %s.amount + $3 WHERE creditor_id=$1 AND debtor_id=$2`, debtsTable, debtsTable)
		_, err = tx.Exec(updateDebtQuery, args...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *PurchasePostgres) Delete(purchase core.Purchase, users []core.UserInputGetAll) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("SELECT amount FROM %s WHERE id=$1", purchasesTable)
	row := tx.QueryRow(query, purchase.ID)
	if err := row.Scan(&purchase.Amount); err != nil {
		tx.Rollback()
		return err
	}

	deletePurchaseQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1", purchasesTable)
	_, err = tx.Exec(deletePurchaseQuery, purchase.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, user := range users {
		if user.Id == purchase.BuyerId {
			continue
		}
		args := make([]interface{}, 0)
		args = append(args, purchase.BuyerId, user.Id, purchase.Amount/float32(len(users)))
		updateDebtQuery := fmt.Sprintf(`UPDATE %s SET amount = %s.amount - $3 WHERE creditor_id=$1 AND debtor_id=$2`, debtsTable, debtsTable)
		_, err = tx.Exec(updateDebtQuery, args...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
