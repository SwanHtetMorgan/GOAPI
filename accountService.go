package main

type AccountService interface {
	CreateAccount(*Account) error
	DeleteAccount(*Account) error
	UpdateAccount(*Account) (*Account, error)
	GetAccountById(id int32) (*Account, error)
	GetAccountByNumber(number int32) (*Account, error)
}
