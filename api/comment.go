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

func GetComment(c *gin.Context) {
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

	mid := c.PostForm("mid")
	if mid == "" {
		log.Printf("mid cannot be empty : %v", err)
		util.RespParamErr(c)
		return
	}
	mID, err := strconv.Atoi(mid)
	if err != nil {
		log.Printf("strconv.Atoi error : %v", err)
		util.RespInternalErr(c)
		return
	}

	msg, err := service.SearchMessageByMID(int64(mID))
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "此MID无对应留言")
		} else {
			log.Printf("GetComment Search Message By MID error : %v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}

	//fmt.Printf("mID : %v\n", mID)
	com, err := service.SearchAllComment(int64(mID))
	//fmt.Printf("com : %v\n", com)

	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "此MID无对应评论")
		} else {
			log.Printf("GetComment Search All Comment error : %v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}

	c.JSON(http.StatusOK, model.AllComment{
		Comment:          msg,
		NextLayerComment: com,
	})
	util.RespOK(c)
	return
}

func CreateComment(c *gin.Context) {
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
	detail := c.PostForm("detail")

	if mid == "" || detail == "" {
		log.Printf("mid and detail cannot be empty : %v", err)
		util.RespParamErr(c)
		return
	}

	mID, err := strconv.Atoi(mid)
	if err != nil {
		log.Printf("MID error : %v", err)
		util.RespParamErr(c)
	}
	MID := int64(mID)

	m, err := service.SearchMessageByMID(MID)

	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "目标留言不存在")
		} else {
			log.Printf("Search Message By MID error : %v", err)
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

	err = service.CreateMessage(model.Message{
		SenderUID: u.ID,
		RecUID:    -1,
		Detail:    detail,
		Thread:    m.Thread + strconv.FormatInt(cid+1, 10) + "/",
	})

	if err != nil {
		fmt.Println("创建评论出问题了")
		fmt.Printf("error : %v\n", err)
		util.RespInternalErr(c)
		return
	}

	util.RespOK(c)
	return
}

func ModifyComment(c *gin.Context) {
	ModifyMessage(c)
	return
}

func DeleteComment(c *gin.Context) {
	DeleteMessage(c)
	return
}
