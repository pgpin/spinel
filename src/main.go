package main

import "github.com/gin-gonic/gin"
import "spinel"
import "encoding/json"
import "io/ioutil"
import "time"

func main() {
	yamlstr, err := ioutil.ReadFile("example-config.yaml")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(gin.Recovery())
//	gin.DisableConsoleColor()
//	f, _ := os.Create("spinel.log")
//	gin.DefaultWriter = io.MultiWriter(f)

	config, _ := spinel.ParseYamlConfiguration(&yamlstr)
	cidrs := spinel.CidrsParse(config.Cidrs)

	//
	// _spinel_auth_check is a route that gets called by Nginx
	// to determine if a request is authenticated. This route should
	// return no content. If the request is allowed this route should
	// return an HTTP 200. If the request is not allowed then it should
	// return an HTTP 401 unauthorized
	//
	r.GET("/_spinel_auth_check", func(c *gin.Context) {
		var token spinel.Token
		//
		// allow request if client ip is in configured whitelists
		//
		if spinel.CidrsContains(&cidrs, c.ClientIP()) {
			c.AbortWithStatus(200)
			return
		}
		//
		// deny request if there is no bearer token cookie
		//
		cookie, err := c.Request.Cookie("spinel_token")
		if err != nil || cookie == nil {
			c.AbortWithStatus(401)
			return
		}
		json.Unmarshal([]byte(cookie.String()), &token)

		//
		// allow request if the bearer token is valid
		// and the token has not expired
		//
		if token.Validate(config.Secret) {
			if time.Now().Unix() > token.Expires {
				c.AbortWithStatus(401)
			} else {
				c.AbortWithStatus(200)
			}
			return
		}

		//
		// default deny
		//
		c.AbortWithStatus(401)
		return
	})
	r.LoadHTMLGlob("tmpl/*")
	r.GET("/_spinel_login", func(c *gin.Context) {
		c.HTML(200, "login.tmpl", gin.H{"url": c.Query("url")})
	})
	r.Static("/_spinel_assets", "./assets") 
	r.Run(config.Listen)
}
