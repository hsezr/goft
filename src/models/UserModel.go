package models

import "fmt"

type UserModel struct {
	UserId int
	UserName string
}

func (this *UserModel) String() string {
	return fmt.Sprintf("userid:%d, username:%s", this.UserId, this.UserName)
}