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
	}

	m := r.Group("/message")
	{
		m.GET("/message", GetMessage)
		m.POST("/message")
		m.PUT("/message")
		m.DELETE("/message")
	}

	r.Run()
}
