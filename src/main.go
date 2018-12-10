package main

import "github.com/gin-gonic/gin"
import "spinel"
import "encoding/json"
import "io/ioutil"
import "time"
import "flag"
import "github.com/jbmcgill/go-throttle"

type LoginPost struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Url      string `form:"url"`
}

func main() {
	configFile := flag.String("file", "/etc/spinel.yaml", "configuration file location")
	flag.Parse()
	yamlstr, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(err)
	}

	config, _ := spinel.ParseYamlConfiguration(&yamlstr)
	cidrs := spinel.CidrsParse(config.Cidrs)
	throttle := &throttle.Throttle{PeriodicityMs: 1000, Limit: config.Ad.MaxRequestsPerSecond}
	ad := spinel.NewActiveDirectoryConnection(config.Ad.Host, config.Ad.Port, config.Ad.Dn)

	if config.Debug != true {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(gin.Recovery())

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
		if spinel.CidrsContains(&cidrs, c.GetHeader("X-Original-IP")) {
			c.AbortWithStatus(200)
			return
		}
		//
		// deny request if there is no bearer token cookie
		//
		cookie, err := c.Cookie("spinel_token")
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		json.Unmarshal([]byte(cookie), &token)

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
		} else {
			c.AbortWithStatus(401)
		}

		//
		// default deny
		//
		c.AbortWithStatus(401)
		return
	})
	r.LoadHTMLGlob(config.Html.Templates + "/*")
	r.GET("/_spinel_login", func(c *gin.Context) {
		c.HTML(200, "login.tmpl", gin.H{"url": c.Query("url"), "loginTitle": config.Html.LoginTitle})
	})
	r.POST("/_spinel_auth", func(c *gin.Context) {
		var login LoginPost
		err := c.ShouldBind(&login)
		if err != nil {
			// bad post
			c.AbortWithStatus(400)
		}
		throttle.Invoke(func() {
			if ad.Authenticate(login.Username, login.Password) {
				token := spinel.NewToken(config.Secret, "*", time.Now().Unix()+60*60*config.Expires)
				c.SetCookie("spinel_token", token.AsJsonString(), 60*60*int(config.Expires), "/", "", false, false)
				c.Redirect(301, login.Url)
			} else {
				// failed to authenticate
				c.AbortWithStatus(401)
			}
		})
	})
	r.RunUnix(config.Socket)
}
