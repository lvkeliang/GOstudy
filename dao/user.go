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
	fmt.Println("执行dao.SearchUserByUserName")
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.ID, &u.UserName, &u.Password)
	return
}

func SearchSecurityQuestionByUID(UID int64) (q model.SecurityQuestion, err error) {
	row := DB.QueryRow("select id,question,answer from security_question where id = ?", UID)
	fmt.Println("执行dao.SearchSecurityQuestionByUID")
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&q.ID, &q.Question, &q.Answer)
	return
}

func ModifyPasswordByUID(UID int64, newpassword string) (err error) {
	fmt.Println("执行dao.ModifyPasswordByUID")
	_, err = DB.Exec("update user set password = ? where id = ?", newpassword, UID)
	return
}

func SetSecurityQuestion(p model.SecurityQuestion) (err error) {
	fmt.Println("执行dao.SetSecurityQuestion")
	_, err = DB.Exec("insert into security_question (id, question,answer) values (?,?,?)", p.ID, p.Question, p.Answer)
	return
}
