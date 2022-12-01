package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"message_board/model"
	"message_board/service"
	"message_board/util"
	"net/http"
	"strconv"
)

func GetAllMessage(c *gin.Context) {
	//mid := c.PostForm("mid")
	cookie, err := c.Cookie("name")
	_, err = service.SearchUserByUserName(cookie)
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

	var AllMsg []model.Message

	i, err := service.GetMessageNumber()
	if err != nil {
		log.Printf("Get Message Number error : %v", err)
		util.RespInternalErr(c)
		return
	}

	for ; i > 0; i-- {
		Msg, err := service.SearchMessageByMID(i)
		if err != nil {
			log.Printf("Search Message By MID error : %v", err)
			util.RespInternalErr(c)
			return
		}
		if Msg.RecUID != -1 {
			if Msg.IsDeleted == 1 {
				AllMsg = append(AllMsg, model.Message{
					MID:       i,
					IsDeleted: 1,
				})
			} else {
				AllMsg = append(AllMsg, Msg)
			}
		}
		//fmt.Printf("i : %v\nAllMsg : %v\n", i, AllMsg)
	}
	c.JSON(http.StatusOK, AllMsg)
	util.RespOK(c)
	return
}

func PostMessage(c *gin.Context) {
	cookie, err := c.Cookie("name")
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

	receiverName := c.PostForm("receiverName")
	detail := c.PostForm("detail")

	if receiverName == "" || detail == "" {
		log.Printf("receiverName and detail cannot be empty : %v\n", err)
		util.RespParamErr(c)
		return
	}

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

	cid, err := service.GetMessageNumber()
	if err != nil {
		log.Printf("Get Message Number error : %v", err)
		util.RespInternalErr(c)
		return
	}

	fmt.Printf("cid: %v \n", cid)
	err = service.CreateMessage(model.Message{
		SenderUID: u.ID,
		RecUID:    receiver.ID,
		Detail:    detail,
		Thread:    "/" + strconv.FormatInt(cid+1, 10) + "/",
	})

	if err != nil {
		fmt.Println("创建留言出问题了")
		util.RespInternalErr(c)
		return
	}

	util.RespOK(c)
	return
}

func ModifyMessage(c *gin.Context) {
	cookie, err := c.Cookie("name")
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

	mid := c.PostForm("mid")

	if mid == "" {
		log.Printf("mid cannot be empty : %v", err)
		util.RespParamErr(c)
		return
	}

	mID, err := strconv.Atoi(mid)
	MID := int64(mID)
	newDetail := c.PostForm("newDetail")

	if newDetail == "" {
		log.Printf("newDetail cannot be empty : %v", err)
		util.RespParamErr(c)
		return
	}

	if err != nil {
		log.Printf("MID receive err : %v", err)
		util.RespInternalErr(c)
		return
	}

	m, err := service.SearchMessageByMID(MID)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "此MID无对应留言")
		} else {
			log.Printf("ModifyMessage Search Message By MID error : %v", err)
			util.RespInternalErr(c)
			return
		}
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
	return
}

func DeleteMessage(c *gin.Context) {
	cookie, err := c.Cookie("name")
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

	mid := c.PostForm("mid")

	if mid == "" {
		log.Printf("mid cannot be empty : %v", err)
		util.RespParamErr(c)
		return
	}

	mID, err := strconv.Atoi(mid)
	MID := int64(mID)

	if err != nil {
		log.Printf("MID receive err : %v", err)
		util.RespInternalErr(c)
		return
	}

	m, err := service.SearchMessageByMID(MID)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "此MID无对应留言")
		} else {
			log.Printf("DeleteMessage Search Message By MID error : %v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}

	if m.SenderUID != u.ID {
		util.NormErr(c, 300, "没有权限")
		return
	}

	if m.IsDeleted == 1 {
		util.NormErr(c, 300, "已经删除")
		return
	}

	err = service.DeleteMessage(MID)
	if err != nil {
		log.Printf("Delete Message error : %v", err)
		util.RespInternalErr(c)
		return
	}

	util.RespOK(c)
	return
}
