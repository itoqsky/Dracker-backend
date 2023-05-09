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
	ID         int     `json:"id" db:"id"`
	CreditorID int     `json:"creditor_id" db:"creditor_id"`
	DebtorID   int     `json:"debtor_id" db:"debtor_id"`
	Amount     float32 `json:"amount" db:"amount"`
}

type Purchase struct {
	ID          int     `json:"id" db:"id"`
	GroupId     int     `json:"group_id" db:"group_id"`
	Amount      float32 `json:"amount" db:"amount" binding:"required"`
	BuyerId     int     `json:"buyer_id" db:"buyer_id"`
	Description string  `json:"description" db:"description" binding:"required"`
	Timestamp   string  `json:"timestamp" db:"timestamp"`
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

type CreatePurchaseResponse struct {
	ID        int    `json:"id" db:"id"`
	Timestamp string `json:"timestamp" db:"timestamp"`
}
