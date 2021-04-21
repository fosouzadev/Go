package models

import "fmt"

type IUser interface {
	ForgotLogin()
}

type User struct {
	IUser   `json:"-"` // json ignore
	Id      int        `json:"id"`
	Login   string     `json:"username"`
	Enabled bool       `json:"-"` // json ignore
}

func (u *User) Init() {
	fmt.Println("construtor")
}

func (u *User) ForgotLogin() {
	fmt.Println("Esqueci a senha")
}
