package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	_const "github.com/yincongcyincong/BasicStudy/library/const"
	"github.com/yincongcyincong/BasicStudy/library/util"
	"github.com/yincongcyincong/BasicStudy/proto/api"
	"github.com/yincongcyincong/BasicStudy/service"
	"net/http"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Main website",
	})
}

func DoLogin(c *gin.Context) {
	input := api.DoLoginReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		logrus.Warningf("doLogin process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	data, errNo, err := service.DoLogin(c, &input)
	if err != nil || errNo != _const.SuccNo {
		logrus.Warningf("doLogin process error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, data)
}
