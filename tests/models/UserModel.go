package models

import "fmt"

type UserModel struct {
	Uid      int
	Username string
}

func (this *UserModel) String() {
	fmt.Println("user model")
}
