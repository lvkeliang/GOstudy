package dao

import (
	"fmt"
	"message_board/model"
)

func InsertUser(u model.User) (err error) {
	fmt.Println("执行InsertUser")
	_, err = DB.Exec("insert into user (name,password) values (?,?)", u.UserName, u.Password)
	return
}

func SearchUserByUserName(name string) (u model.User, err error) {
	row := DB.QueryRow("select id,name,password from user where name = ?", name)
	fmt.Println("执行SearchUserByUserName")
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.ID, &u.UserName, &u.Password)
	return
}
