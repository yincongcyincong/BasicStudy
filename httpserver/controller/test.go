package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	_const "github.com/yincongcyincong/BasicStudy/library/const"
	"github.com/yincongcyincong/BasicStudy/library/util"
	"github.com/yincongcyincong/BasicStudy/service"

	"github.com/yincongcyincong/BasicStudy/proto/api"
)

// addStudyTest 新增信息
func addStudyTest(c *gin.Context) {
	input := api.AddStudyTestReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		glog.Warningf("addStudyTest input process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	data, errNo, err := service.AddStudyTest(c, &input)
	if err != nil || errNo != _const.SuccNo {
		glog.Warningf("addStudyTest error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, data)
}

// updateStudyTest 修改信息

func updateStudyTest(c *gin.Context) {
	input := api.UpdateStudyTestReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		glog.Warningf("updateStudyTest input process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	errNo, err := service.UpdateStudyTest(c, &input)
	if err != nil || errNo != _const.SuccNo {
		glog.Warningf("updateStudyTest error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, nil)

}

// delStudyTest 删除信息

func delStudyTest(c *gin.Context) {
	input := api.DelStudyTestReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		glog.Warningf("delStudyTest input process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	errNo, err := service.DelStudyTest(c, &input)
	if err != nil || errNo != _const.SuccNo {
		glog.Warningf("delStudyTest error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, nil)
}

// mgetStudyTestByIDs 按照ID批量查询信息

func mgetStudyTestByIDs(c *gin.Context) {
	input := api.MgetStudyTestByIDsReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		glog.Warningf("mgetStudyTestByIDs input process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	data, errNo, err := service.MgetStudyTestByIDs(c, &input)
	if err != nil || errNo != _const.SuccNo {
		glog.Warningf("mgetStudyTestByIDs error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, data)
}

// mgetStudyTestByCond 按照条件查询信息（无缓存）

func mgetStudyTestByCond(c *gin.Context) {
	input := api.MgetStudyTestByCondReq{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		glog.Warningf("mgetStudyTestByCond input process error, err=%v", err)
		util.FailureWithErrMsg(c, _const.ParamErrorNo, "")
		return
	}

	data, errNo, err := service.MgetStudyTestByCond(c, &input)
	if err != nil || errNo != _const.SuccNo {
		glog.Warningf("mgetStudyTestByCond error, err=%v", err)
		util.FailureWithErrMsg(c, errNo, "")
		return
	}
	util.Success(c, data)
}
