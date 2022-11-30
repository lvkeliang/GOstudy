package dao

import (
	"fmt"
	"message_board/model"
)

func InsertMessage(m model.Message) (err error) {
	fmt.Println("执行dao.InsertMessage")
	_, err = DB.Exec("insert into message (send_uid,rec_uid,detail) values (?,?,?)", m.SenderUID, m.RecUID, m.Detail)
	return
}

func SearchMessageByMID(MID int64) (m model.Message, err error) {
	fmt.Println("执行dao.SearchUserByUserName")
	row := DB.QueryRow("select mid,send_uid,rec_uid,detail,is_modified,is_deleted from message where mid = ?", MID)

	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&m.MID, &m.SenderUID, &m.RecUID, &m.Detail, &m.IsModified, &m.IsDeleted)
	return
}

func ModifyMessage(MID int64, newDetail string) (err error) {
	fmt.Println("执行dao.ModifyMessage")
	_, err = DB.Exec("update message set detail = ? where mid = ?", newDetail, MID)
	_, err = DB.Exec("update message set is_modified = ? where mid = ?", 1, MID)
	return
}

func DeleteMessage(MID int64) (err error) {
	fmt.Println("执行dao.DeleteMessage")
	_, err = DB.Exec("update message set is_deleted = ? where mid = ?", 1, MID)
	return
}
