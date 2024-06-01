package util

import (
	"github.com/gin-gonic/gin"
	_const "github.com/yincongcyincong/BasicStudy/library/const"
	"net/http"
)

// util.FailureWithErrMsg 指定错误号和错误信息的返回
func FailureWithErrMsg(ctx *gin.Context, errno int, errmsg string) {
	if len(errmsg) == 0 {
		errmsg = _const.GetErrorMsg(errno)
	}

	response := gin.H{
		"errno":  errno,
		"errmsg": errmsg,
	}

	ctx.JSON(http.StatusOK, response)
	return
}

func Success(ctx *gin.Context, data interface{}) {
	response := gin.H{
		"errno":  _const.SuccMsg,
		"errmsg": _const.SuccMsg,
		"data":   data,
	}
	ctx.JSON(http.StatusOK, response)
	return
}
