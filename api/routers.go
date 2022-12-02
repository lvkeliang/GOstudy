package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()

	u := r.Group("/user")
	{
		u.GET("/register", func(c *gin.Context) {
			c.String(200, "ok\n")
			return
		})
		u.POST("/register", Register)
		u.POST("/login", Login)
		u.POST("/securityQuestion", SetSecurityQuestion)
		u.PUT("/resetPassword", ResetPassword)
		u.DELETE("/announcement", Announcement)
	}

	m := r.Group("/message")
	{
		m.GET("/message", GetAllMessage)
		m.POST("/message", PostMessage)
		m.PUT("/message", ModifyMessage)
		m.DELETE("/message", DeleteMessage)

		co := m.Group("/comment")
		{
			co.GET("/comment", GetComment)
			co.POST("/comment", CreateComment)
			co.PUT("/comment", ModifyComment)
			co.DELETE("/comment", DeleteComment)
		}
	}

	r.Run()
}
