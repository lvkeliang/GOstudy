package service

import (
	"fmt"
	"message_board/dao"
	"message_board/model"
)

func SearchUserByUserName(name string) (u model.User, err error) {
	fmt.Println("执行SearchUserByUserName")
	u, err = dao.SearchUserByUserName(name)
	return
}

func CreateUser(u model.User) error {
	fmt.Println("执行CreateUser")
	err := dao.InsertUser(u)
	return err
}
