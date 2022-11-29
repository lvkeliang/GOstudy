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
