package core

type Group struct {
	ID        int        `json:"-"`
	Name      string     `json:"name"`
	Members   []UserID   `json:"members"`
	Debts     []Debt     `json:"debts"`
	Purchases []Purchase `json:"purchases"`
}

type UserID int

type Debt struct {
	ID       int    `json:"id"`
	Debtor   UserID `json:"debtor"`
	Creditor UserID `json:"creditor"`
	Amount   int    `json:"amount"`
}

type Purchase struct {
	ID          int    `json:"id"`
	SpentBy     UserID `json:"spent_by"`
	Amount      int    `json:"amount"`
	Date        string `json:"date"`
	Description string `json:"description"`
}
