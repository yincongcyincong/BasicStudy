package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yincongcyincong/BasicStudy/bootstrap"
	_const "github.com/yincongcyincong/BasicStudy/library/const"
	"github.com/yincongcyincong/BasicStudy/library/util"
	"github.com/yincongcyincong/BasicStudy/model/dao"
	"github.com/yincongcyincong/BasicStudy/proto/api"
	"google.golang.org/protobuf/proto"
	"strings"
)

// AddStudyQuestion 新增信息

func AddStudyQuestion(c *gin.Context, input *api.AddStudyQuestionReq) (*api.AddStudyQuestionData, int, error) {
	// 参数校验
	if input == nil {
		logrus.Warningf("AddStudyQuestion input invalid, input=%+v", input)
		return nil, _const.ParamErrorNo, errors.New("input params invalid")
	}

	//message := []string{
	//	fmt.Sprintf(_const.CNCalPrompt, 10, 100, "加减乘除"),
	//}
	//questionInfo, err := util.GPT(bootstrap.GPTConfigInstance.SecretKey, message)

	message := []string{
		fmt.Sprintf(_const.ErNiePrompt, 10, 100, "加减乘除"),
	}
	questionInfo, err := util.ErNie(bootstrap.ErNieConfigInstance.Ak, bootstrap.ErNieConfigInstance.Sk, message)
	fmt.Println(questionInfo)

	// 插入DB
	db := dao.NewStudyQuestionDao()
	insertReq := dao.StudyQuestionItem{}

	if input.CreateTime != nil {
		insertReq.CreateTime = proto.Uint64(input.GetCreateTime())
	}

	qa := strings.Split(questionInfo, "参考答案")
	insertReq.Question = proto.String(qa[0])
	insertReq.Answer = proto.String(qa[1])
	if input.TestId != nil {
		insertReq.TestId = proto.Uint64(input.GetTestId())
	}
	if input.Type != nil {
		infoType := uint8(input.GetType())
		insertReq.Type = &infoType
	}

	id, err := db.Add(context.Background(), insertReq)
	if err != nil {
		logrus.Warningf("AddStudyQuestion insert db fail, err=%v, input=%+v", err, insertReq)
		return nil, _const.DBExecError, err
	}

	data := &api.AddStudyQuestionData{
		Id: proto.Uint64(uint64(id)),
	}
	return data, _const.SuccNo, nil
}

// UpdateStudyQuestion 修改信息

func UpdateStudyQuestion(c *gin.Context, input *api.UpdateStudyQuestionReq) (int, error) {
	// 参数校验
	if input == nil || input.GetId() <= 0 {
		logrus.Warningf("UpdateStudyQuestion input invalid, input=%+v", input)
		return _const.ParamErrorNo, errors.New("input params invalid")
	}

	id := input.GetId()

	// 更新DB
	db := dao.NewStudyQuestionDao()
	updateReq := dao.StudyQuestionItem{}

	if input.CreateTime != nil {
		updateReq.CreateTime = proto.Uint64(input.GetCreateTime())
	}
	if input.TestId != nil {
		updateReq.TestId = proto.Uint64(input.GetTestId())
	}
	if input.Type != nil {
		infoType := uint8(input.GetType())
		updateReq.Type = &infoType
	}

	_, err := db.Update(context.Background(), updateReq, id)
	if err != nil {
		logrus.Warningf("UpdateStudyQuestion update db fail, err=%v, input=%+v", err, updateReq)
		return _const.DBQueryError, err
	}
	return _const.SuccNo, nil
}

// DelStudyQuestion 删除信息

func DelStudyQuestion(c *gin.Context, input *api.DelStudyQuestionReq) (int, error) {
	// 参数校验
	if input == nil || input.GetId() <= 0 {
		logrus.Warningf("DelStudyQuestion input invalid, input=%+v", input)
		return _const.ParamErrorNo, errors.New("input params invalid")
	}

	id := input.GetId()
	// 更新DB
	db := dao.NewStudyQuestionDao()
	delReq := dao.StudyQuestionItem{}

	_, err := db.Update(context.Background(), delReq, id)
	if err != nil {
		logrus.Warningf("DelStudyQuestion update db fail, err=%v, input=%+v", err, delReq)
		return _const.DBQueryError, err
	}
	return _const.SuccNo, nil
}

// MgetStudyQuestionByIDs 按照ID批量查询信息

func MgetStudyQuestionByIDs(c *gin.Context, input *api.MgetStudyQuestionByIDsReq) (*api.MgetStudyQuestionByIDsData, int, error) {
	// 参数校验

	if input == nil || len(input.GetIds()) <= 0 {
		logrus.Warningf("MgetStudyQuestionByIDs input invalid, input=%+v", input)
		return nil, _const.ParamErrorNo, errors.New("input params invalid")
	}

	// 待查询数据去重
	queryIDs := make([]uint64, 0, len(input.GetIds()))
	hits := map[uint64]struct{}{}
	for _, id := range input.GetIds() {
		if _, ok := hits[id]; ok {
			continue
		}
		queryIDs = append(queryIDs, id)
		hits[id] = struct{}{}
	}

	info := map[uint64]*api.StudyQuestion{}

	// 查询DB
	db := dao.NewStudyQuestionDao()
	queryRes, err := db.MgetByIDs(context.Background(), queryIDs)
	if err != nil {
		logrus.Warningf("MgetStudyQuestionByIDs query db fail, err=%v, query_id=%+v", err, &queryIDs)
		return nil, _const.DBQueryError, err
	}
	for _, res := range formatStudyQuestion(queryRes) {
		info[res.GetId()] = res
	}

	data := &api.MgetStudyQuestionByIDsData{
		Info: info,
	}
	return data, _const.SuccNo, nil
}

// MgetStudyQuestionByCond 按照条件查询信息（无缓存）

func MgetStudyQuestionByCond(c *gin.Context, input *api.MgetStudyQuestionByCondReq) (*api.MgetStudyQuestionByCondData, int, error) {
	// 参数校验

	if input == nil {
		logrus.Warningf("MgetStudyQuestionByCond input invalid, input=%+v", input)
		return nil, _const.ParamErrorNo, errors.New("input params invalid")
	}

	// 查询DB
	db := dao.NewStudyQuestionDao()
	queryReq := dao.StudyQuestionItem{}

	queryRes, totalNum, err := db.MgetByPage(context.Background(), queryReq, input.GetPn(), input.GetRn())
	if err != nil {
		logrus.Warningf("MgetStudyQuestionByCond query db fail, err=%v, input=%+v", err, input)
		return nil, _const.DBQueryError, err
	}

	data := &api.MgetStudyQuestionByCondData{
		Info:     formatStudyQuestion(queryRes),
		TotalNum: proto.Uint32(uint32(totalNum)),
	}
	return data, _const.SuccNo, nil
}

// formatStudyQuestion 格式化返回值，将内部使用的dao结构体转换为对外使用的proto结构体
func formatStudyQuestion(srcData []dao.StudyQuestionItem) []*api.StudyQuestion {
	dstData := make([]*api.StudyQuestion, 0, len(srcData))
	for _, data := range srcData {
		info := api.StudyQuestion{}
		if data.Id != nil {
			info.Id = proto.Uint64(*data.Id)
		}
		if data.CreateTime != nil {
			info.CreateTime = proto.Uint64(*data.CreateTime)
		}
		if data.Question != nil {
			info.Question = data.Question
		}
		if data.Answer != nil {
			info.Answer = data.Answer
		}
		if data.TestId != nil {
			info.TestId = proto.Uint64(*data.TestId)
		}
		if data.Type != nil {
			infoType := uint32(*data.Type)
			info.Type = proto.Uint32(infoType)
		}
		dstData = append(dstData, &info)
	}
	return dstData
}
