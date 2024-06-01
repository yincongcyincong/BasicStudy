package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	_const "github.com/yincongcyincong/BasicStudy/library/const"
	"github.com/yincongcyincong/BasicStudy/library/util"
	"github.com/yincongcyincong/BasicStudy/service"

	"github.com/yincongcyincong/BasicStudy/proto/api"
)

// addStudyQuestion 新增信息
func AddStudyQuestion(c *gin.Context) {
	input := api.AddStudyQuestionReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	data, errNo, err := service.AddStudyQuestion(c, &input)
	if err != nil || errNo != _const.SuccNo {
		glog.Warningf("addStudyQuestion process error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, data)
}

// updateStudyQuestion 修改信息
func updateStudyQuestion(c *gin.Context) {
	input := api.UpdateStudyQuestionReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		glog.Warningf("updateStudyQuestion input process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	errNo, err := service.UpdateStudyQuestion(c, &input)
	if err != nil || errNo != _const.SuccNo {
		glog.Warningf("updateStudyQuestion process error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, nil)
}

// delStudyQuestion 删除信息
func delStudyQuestion(c *gin.Context) {
	input := api.DelStudyQuestionReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		glog.Warningf("delStudyQuestion input process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	errNo, err := service.DelStudyQuestion(c, &input)
	if err != nil || errNo != _const.SuccNo {
		glog.Warningf("delStudyQuestion process error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, nil)
}

// mgetStudyQuestionByIDs 按照ID批量查询信息
func mgetStudyQuestionByIDs(c *gin.Context) {
	input := api.MgetStudyQuestionByIDsReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		glog.Warningf("mgetStudyQuestionByIDs input process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	data, errNo, err := service.MgetStudyQuestionByIDs(c, &input)
	if err != nil || errNo != _const.SuccNo {
		glog.Warningf("mgetStudyQuestionByIDs process error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, data)
}

// mgetStudyQuestionByCond 按照条件查询信息（无缓存）
func mgetStudyQuestionByCond(c *gin.Context) {
	input := api.MgetStudyQuestionByCondReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		glog.Warningf("mgetStudyQuestionByCond input process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	data, errNo, err := service.MgetStudyQuestionByCond(c, &input)
	if err != nil || errNo != _const.SuccNo {
		glog.Warningf("mgetStudyQuestionByCond process error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, data)
}
