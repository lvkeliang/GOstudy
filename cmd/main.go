package main

import (
	"message_board/api"
	"message_board/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
