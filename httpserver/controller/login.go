package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yincongcyincong/BasicStudy/library/util"
)

func ChooseUser(c *gin.Context) {
	u := c.Query("username")
	if u == "" {
		u = "basic_study"
	}
	cookie, err := c.Cookie("username")
	if err != nil || cookie == "" {
		c.SetCookie("username", u, 3600000, "/", c.Request.Host, false, true)
	}

	util.Success(c, nil)
}


