package storage

import (
	"fmt"

	"github.com/itoqsky/money-tracker-backend/internal/core"
	"github.com/jmoiron/sqlx"
)

type DebtPostgres struct {
	db *sqlx.DB
}

func NewDebtPostgres(db *sqlx.DB) *DebtPostgres {
	return &DebtPostgres{db: db}
}

func (s *DebtPostgres) GetAll(userId int) ([]core.Debt, []core.Debt, error) {
	var debts []core.Debt
	var credits []core.Debt

	tx, err := s.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	stmt, err := tx.Prepare(`SELECT d.creditor_id, d.debtor_id, d.amount FROM debts d WHERE d.debtor_id=$1`)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var debt core.Debt
		err := rows.Scan(&debt.CreditorID, &debt.DebtorID, &debt.Amount)
		if err != nil {
			tx.Rollback()
			return nil, nil, err
		}
		debts = append(debts, debt)
	}

	err = rows.Err()
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	stmt, err = tx.Prepare(`SELECT d.creditor_id, d.debtor_id, d.amount FROM debts d WHERE d.creditor_id=$1`)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	defer stmt.Close()

	rows, err = stmt.Query(userId)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var debt core.Debt
		err := rows.Scan(&debt.CreditorID, &debt.DebtorID, &debt.Amount)
		if err != nil {
			tx.Rollback()
			return nil, nil, err
		}
		credits = append(credits, debt)
	}

	err = rows.Err()
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	return debts, credits, tx.Commit()
}

func (s *DebtPostgres) Update(debt core.Debt) error {
	query := fmt.Sprintf(`INSERT INTO %s (creditor_id, debtor_id, amount) VALUES ($1, $2, $3) 
		ON CONFLICT (creditor_id, debtor_id) DO UPDATE SET amount = %s.amount + $3`, debtsTable, debtsTable)
	_, err := s.db.Exec(query, debt.CreditorID, debt.DebtorID, debt.Amount)
	return err
}
