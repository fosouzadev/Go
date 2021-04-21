package models

import "fmt"

type User struct {
	Id    int
	Login string
}

func (u *User) Init() {
	fmt.Println("construtor")
}
