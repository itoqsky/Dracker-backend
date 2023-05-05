package storage

type Authorization interface {
}

type Debt interface {
}

type Purchase interface {
}

type User interface {
}

type Storage struct {
	Authorization
	Debt
	Purchase
	User
}

func NewStorage() *Storage {
	return &Storage{}
}
