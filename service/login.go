package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	_const "github.com/yincongcyincong/BasicStudy/library/const"
	"github.com/yincongcyincong/BasicStudy/library/util"
	"github.com/yincongcyincong/BasicStudy/model/dao"
	"github.com/yincongcyincong/BasicStudy/proto/api"
)

func DoLogin(c *gin.Context, input *api.DoLoginReq) (*api.AddStudyQuestionData, int, error) {
	db := dao.NewStudyUserDao()
	queryRes, err := db.MgetByUsername(c, input.GetUsername())
	if err != nil {
		logrus.Warning("dologin fail: %v", err)
		return nil, _const.DBQueryError, err
	}

	if queryRes == nil || queryRes.Password == nil {
		logrus.Warning("username not exist")
		return nil, _const.AuthError, err
	}

	if *queryRes.Password != util.Md5(input.GetPassword()) {
		logrus.Warning("password not match")
		return nil, _const.AuthError, err
	}
	return nil, _const.SuccNo, nil
}
