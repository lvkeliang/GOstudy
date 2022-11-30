package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"message_board/model"
	"message_board/service"
	"message_board/util"
	"strconv"
)

func GetMessage(c *gin.Context) {
	//mid := c.PostForm("mid")
}

func PostMessage(c *gin.Context) {
	u, err := service.IsLoggedIn(c)
	if err != nil {
		log.Printf("Is Logged In error : %v", err)
		util.RespInternalErr(c)
		return
	}

	if u.UserName == "" {
		util.NormErr(c, 300, "未登录")
		return
	}

	receiverName := c.PostForm("receiverName")
	detail := c.PostForm("detail")

	receiver, err := service.SearchUserByUserName(receiverName)

	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "目标用户不存在")
		} else {
			log.Printf("search user error : %v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}

	err = service.CreateMessage(model.Message{
		SenderUID: u.ID,
		RecUID:    receiver.ID,
		Detail:    detail,
	})

	if err != nil {
		fmt.Println("创建留言出问题了")
		util.RespInternalErr(c)
		return
	}

	util.RespOK(c)
}

func ModifyMessage(c *gin.Context) {
	u, err := service.IsLoggedIn(c)
	if err != nil {
		log.Printf("Is Logged In error : %v", err)
		util.RespInternalErr(c)
		return
	}

	if u.UserName == "" {
		util.NormErr(c, 300, "未登录")
		return
	}

	mID, err := strconv.Atoi(c.PostForm("mid"))
	MID := int64(mID)
	newDetail := c.PostForm("newDetail")

	if err != nil {
		log.Printf("MID receive err : %v", err)
		util.RespInternalErr(c)
		return
	}

	m, err := service.SearchMessageByMID(MID)
	if err != nil {
		log.Printf("Search Message By MID error : %v", err)
		util.RespInternalErr(c)
		return
	}

	if m.SenderUID != u.ID {
		util.NormErr(c, 300, "没有权限")
		return
	}

	if newDetail == "" {
		fmt.Println("输入空字符")
		util.RespParamErr(c)
		return
	}

	err = service.ModifyMessage(MID, newDetail)
	if err != nil {
		log.Printf("Search Message By MID error : %v", err)
		util.RespInternalErr(c)
		return
	}

	util.RespOK(c)
}

func DeleteMessage(c *gin.Context) {
	u, err := service.IsLoggedIn(c)
	if err != nil {
		log.Printf("Is Logged In error : %v", err)
		util.RespInternalErr(c)
		return
	}

	if u.UserName == "" {
		util.NormErr(c, 300, "未登录")
		return
	}

	mID, err := strconv.Atoi(c.PostForm("mid"))
	MID := int64(mID)

	if err != nil {
		log.Printf("MID receive err : %v", err)
		util.RespInternalErr(c)
		return
	}

	m, err := service.SearchMessageByMID(MID)
	if err != nil {
		log.Printf("Search Message By MID error : %v", err)
		util.RespInternalErr(c)
		return
	}

	if m.SenderUID != u.ID {
		util.NormErr(c, 300, "没有权限")
		return
	}

	err = service.DeleteMessage(MID)
	if err != nil {
		log.Printf("Delete Message error : %v", err)
		util.RespInternalErr(c)
		return
	}

	util.RespOK(c)
}
