package service

import (
	"fmt"
	"message_board/dao"
	"message_board/model"
)

func SearchAllComment(MID int64) (commentlist map[int64]model.AllComment, err error) {
	fmt.Println("执行service.SearchAllComment")
	commentlist, err = dao.SearchAllComment(MID)
	return
}
