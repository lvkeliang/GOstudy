package service

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"message_board/dao"
	"message_board/model"
	"message_board/util"
)

func SearchUserByUserName(name string) (u model.User, err error) {
	fmt.Println("执行service.SearchUserByUserName")
	u, err = dao.SearchUserByUserName(name)
	return
}

func CreateUser(u model.User) error {
	fmt.Println("执行CreateUser")
	err := dao.InsertUser(u)
	return err
}

func SearchSecurityQuestionByUID(UID int64) (u model.SecurityQuestion, err error) {
	fmt.Println("执行service.SearchSecurityQuestionByUID")
	u, err = dao.SearchSecurityQuestionByUID(UID)
	return
}

func ModifyPasswordByUID(UID int64, newpassword string) (err error) {
	fmt.Println("执行service.ModifyPasswordByUID")
	err = dao.ModifyPasswordByUID(UID, newpassword)
	return
}

func SetSecurityQuestion(p model.SecurityQuestion) (err error) {
	fmt.Println("执行service.SetSecurityQuestion")
	err = dao.SetSecurityQuestion(p)
	return
}

func IsLoggedIn(c *gin.Context) (u model.User, err error) {
	cookie, err := c.Cookie("name")
	u, err = SearchUserByUserName(cookie)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "未登录")
		} else {
			log.Printf("Is Logged In error : %v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	return
}
