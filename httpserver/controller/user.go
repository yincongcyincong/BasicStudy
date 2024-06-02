package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	_const "github.com/yincongcyincong/BasicStudy/library/const"
	"github.com/yincongcyincong/BasicStudy/library/util"
	"github.com/yincongcyincong/BasicStudy/proto/api"
	"github.com/yincongcyincong/BasicStudy/service"
)

// addStudyUser

func AddStudyUser(c *gin.Context) {
	input := api.AddStudyUserReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		logrus.Warningf("addStudyUser input process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	data, errNo, err := service.AddStudyUser(c, &input)
	if err != nil || errNo != _const.SuccNo {
		logrus.Warningf("addStudyUser error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, data)
}

// updateStudyUser 修改信息

func UpdateStudyUser(c *gin.Context) {
	input := api.UpdateStudyUserReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		logrus.Warningf("updateStudyUser input process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	errNo, err := service.UpdateStudyUser(c, &input)
	if err != nil || errNo != _const.SuccNo {
		logrus.Warningf("updateStudyUser error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, nil)
}

// delStudyUser 删除信息

func DelStudyUser(c *gin.Context) {
	input := api.DelStudyUserReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		logrus.Warningf("delStudyUser input process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	errNo, err := service.DelStudyUser(c, &input)
	if err != nil || errNo != _const.SuccNo {
		logrus.Warningf("delStudyUser error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, nil)
}

// mgetStudyUserByIDs 按照ID批量查询信息

func MgetStudyUserByIDs(c *gin.Context) {
	input := api.MgetStudyUserByIDsReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		logrus.Warningf("mgetStudyUserByIDs input process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	data, errNo, err := service.MgetStudyUserByIDs(c, &input)
	if err != nil || errNo != _const.SuccNo {
		logrus.Warningf("mgetStudyUserByIDs error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, data)
}

// mgetStudyUserByCond 按照条件查询信息（无缓存）

func MgetStudyUserByCond(c *gin.Context) {
	input := api.MgetStudyUserByCondReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		logrus.Warningf("mgetStudyUserByCond input process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	data, errNo, err := service.MgetStudyUserByCond(c, &input)
	if err != nil || errNo != _const.SuccNo {
		logrus.Warningf("mgetStudyUserByCond error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, data)
}
