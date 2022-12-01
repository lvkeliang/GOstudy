package dao

import (
	"fmt"
	"message_board/model"
)

func InsertMessage(m model.Message) (err error) {
	fmt.Println("执行dao.InsertMessage")
	_, err = DB.Exec("insert into message (send_uid,rec_uid,detail,thread) values (?,?,?,?)", m.SenderUID, m.RecUID, m.Detail, m.Thread)
	return
}

func SearchMessageByMID(MID int64) (m model.Message, err error) {
	fmt.Println("执行dao.SearchMessageByMID")
	row := DB.QueryRow("select mid,send_uid,rec_uid,detail,is_modified,is_deleted,thread from message where mid = ?", MID)

	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&m.MID, &m.SenderUID, &m.RecUID, &m.Detail, &m.IsModified, &m.IsDeleted, &m.Thread)
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

func GetMessageNumber() (num int64, err error) {
	num = 0
	fmt.Println("执行dao.GetMessageNumber")
	row, err := DB.Query("select count(*) from message")
	if err = row.Err(); row.Err() != nil {
		return
	}
	for row.Next() {
		err = row.Scan(&num)
		if err != nil {
			return
		}
	}
	return
}
