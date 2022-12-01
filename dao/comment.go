package dao

import (
	"database/sql"
	"fmt"
	"message_board/model"
	"strconv"
)

func SearchAllComment(MID int64) (commentlist map[int64]model.AllComment, err error) {
	commentlist = make(map[int64]model.AllComment)
	fmt.Println("执行dao.SearchAllComment")
	m, err := SearchMessageByMID(MID)
	com := model.Message{}
	if err != nil {
		return nil, err
	}

	for i, err := GetMessageNumber(); i > 0; i-- {
		//fmt.Println("运行dao.GetMessageNumber")
		//fmt.Printf("i : %v\n", i)
		if err != nil {
			return nil, err
		}

		//fmt.Printf("thread : %v\n", m.Thread+strconv.FormatInt(i, 10)+"/")
		row := DB.QueryRow("select mid, send_uid, rec_uid, detail, is_modified, is_deleted, thread from message where thread = ?", m.Thread+strconv.FormatInt(i, 10)+"/")
		err = row.Scan(&com.MID, &com.SenderUID, &com.RecUID, &com.Detail, &com.IsModified, &com.IsDeleted, &com.Thread)

		if err != nil {
			//fmt.Printf("err : %v\n", err)
			if err == sql.ErrNoRows || i != 0 {
				err = nil
				continue
			} else if err == sql.ErrNoRows || i == 0 {
				fmt.Printf("dao.com : %v\n", com)
				return commentlist, nil
			} else {
				return nil, err
			}
		}

		if com.IsDeleted == 1 {
			nextLayerComment, err := SearchAllComment(com.MID)
			if err != nil {
				return nil, err
			}
			commentlist[i] = model.AllComment{
				Comment: model.Message{
					MID:       i,
					IsDeleted: 1,
					Thread:    com.Thread,
				},
				NextLayerComment: nextLayerComment,
			}
		} else {
			nextLayerComment, err := SearchAllComment(com.MID)
			if err != nil {
				return nil, err
			}
			commentlist[i] = model.AllComment{
				Comment:          com,
				NextLayerComment: nextLayerComment,
			}
		}
	}
	return
}
