package main

import "github.com/gin-gonic/gin"
import "spinel"
import "encoding/json"

func main() {
	secret := "changeme"
	r := gin.Default()
	r.GET("/_spinel_auth_check", func(c *gin.Context) {
		var token spinel.Token
		cookie, err := c.Request.Cookie("spinel_token")
		if err != nil || cookie == nil {
			c.AbortWithStatus(401)
			return
		}
		json.Unmarshal([]byte(cookie.String()), &token)

		if !token.Validate(secret) {
			c.AbortWithStatus(401)
			return
		}
		c.AbortWithStatus(200)
	})
	r.GET("/_spinel_login", func(c *gin.Context) { c.JSON(200, gin.H{"foo": "bar"}) })
	r.Run()
}
