package core

import "errors"

type Group struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}

type UsersGroup struct {
	ID      int `json:"id"`
	UserID  int `json:"user_id"`
	GroupID int `json:"group_id"`
}

type Debt struct {
	ID         int     `json:"id"`
	DebtorID   int     `json:"debtor_id"`
	CreditorID int     `json:"creditor_id"`
	Amount     float32 `json:"amount"`
}

type Purchase struct {
	ID          int     `json:"id"`
	SpentByID   int     `json:"spent_by_id"`
	Amount      float32 `json:"amount"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
}

type UpdateGroupInput struct {
	Name *string `json:"name"`
}

func (i *UpdateGroupInput) Validate() error {
	if i.Name == nil {
		return errors.New("update payload has no required fields")
	}
	return nil
}
