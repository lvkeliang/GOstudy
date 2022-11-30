package service

import (
	"fmt"
	"message_board/dao"
	"message_board/model"
)

func CreateMessage(m model.Message) (err error) {
	fmt.Println("执行service.CreateMessage")
	err = dao.InsertMessage(m)
	return err
}

func SearchMessageByMID(MID int64) (m model.Message, err error) {
	fmt.Println("执行service.SearchMessageByUID")
	m, err = dao.SearchMessageByMID(MID)
	return
}

func ModifyMessage(MID int64, newDetail string) (err error) {
	fmt.Println("执行service.ModifyMessage")
	err = dao.ModifyMessage(MID, newDetail)
	return
}

func DeleteMessage(MID int64) (err error) {
	fmt.Println("执行service.DeleteMessage")
	err = dao.DeleteMessage(MID)
	return
}
