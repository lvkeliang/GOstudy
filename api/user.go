package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"message_board/model"
	"message_board/service"
	"message_board/util"
)

func Register(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password")
	fmt.Printf("username: %v\npassword: %v\n", username, password)
	if username == "" || password == "" {
		fmt.Println("输入空字符")
		util.RespParamErr(c)
		return
	}

	u, err := service.SearchUserByUserName(username)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("service.SearchUserByUserName出问题了")
		log.Printf("search user error : %v", err)
		util.RespInternalErr(c)
		return
	}

	if u.UserName != "" {
		util.NormErr(c, 300, "账户已存在")
		return
	}

	err = service.CreateUser(model.User{
		UserName: username,
		Password: password,
	})

	if err != nil {
		fmt.Println("创建用户出问题了")
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func Login(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password")
	fmt.Printf("username: %v\npassword: %v\n", username, password)
	if username == "" || password == "" {
		fmt.Println("输入空字符")
		util.RespParamErr(c)
		return
	}

	u, err := service.SearchUserByUserName(username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user error : %v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	if u.Password != password {
		util.NormErr(c, 20001, "密码错误")
		return
	}

	c.SetCookie("name", username, 0, "", "/", false, false)
	util.RespOK(c)
	fmt.Println("Cookie设置成功")
}

func SetSecurityQuestion(c *gin.Context) {
	cookie, err := c.Cookie("name")
	//fmt.Printf("cookie: %v\nquestion: %v\nanswer: %v\n", cookie, question, answer)
	u, err := service.SearchUserByUserName(cookie)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "未登录")
		} else {
			log.Printf("ResetPassword search user error : %v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}

	q, err := service.SearchSecurityQuestionByUID(u.ID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("search security question error : %v", err)
			util.RespInternalErr(c)
			return
		}
	}

	if q.Question != "" {
		util.NormErr(c, 20003, "已经设置过密保了")
		return
	}

	question := c.PostForm("question")
	answer := c.PostForm("answer")

	if question == "" || answer == "" {
		fmt.Println("输入空字符")
		util.RespParamErr(c)
		return
	}

	err = service.SetSecurityQuestion(model.SecurityQuestion{
		ID:       u.ID,
		Question: question,
		Answer:   answer,
	})

	if err != nil {
		log.Printf("Set Security Question error : %v", err)
		util.RespInternalErr(c)
		return
	}

	util.RespOK(c)
	fmt.Println("密保设置成功")
}

func ResetPassword(c *gin.Context) {
	cookie, err := c.Cookie("name")
	answer := c.PostForm("answer")
	newPassword := c.PostForm("newpassword")
	fmt.Printf("cookie: %v\nanswer: %v\nnewPassword: %v\n", cookie, answer, newPassword)
	u, err := service.SearchUserByUserName(cookie)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "未登录")
		} else {
			log.Printf("ResetPassword search user error : %v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}

	if answer == "" {
		fmt.Println("输入空字符")
		util.RespParamErr(c)
		return
	}

	q, err := service.SearchSecurityQuestionByUID(u.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "未设置密保问题")
		} else {
			log.Printf("search securityquestion error : %v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}

	if q.Answer != answer {
		util.NormErr(c, 20002, "密保问题回答错误")
		return
	}

	if newPassword == "" {
		fmt.Println("输入空字符")
		util.RespParamErr(c)
		return
	}

	err = service.ModifyPasswordByUID(u.ID, newPassword)
	if err != nil {
		log.Printf("Modify Password error : %v", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
	fmt.Println("密码重设成功")
}
