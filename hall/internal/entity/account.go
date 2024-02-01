// @Author huzejun 2024/2/1 17:43:00
package entity

import "github.com/thkhxm/tgf/db"

type UserAccount struct {
	db.Model
	Account  string `orm:"pk"`
	Password string `orm:"pk"`
	UserId   string
}

func (u *UserAccount) GetTableName() string {
	return "t_account"
}

type UserData struct {
	UserId string `orm:"pk"`
	Name   string
}
