package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// not working from Postman, try adding a UI here

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	r.GET("/protected", func(c *gin.Context) {
		c.String(200, csrf.GetToken(c))
	})

	r.POST("/protected", func(c *gin.Context) {
		c.String(200, "CSRF token is valid")
	})

	r.Run(":8080")
}
