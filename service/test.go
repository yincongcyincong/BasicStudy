package service

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	_const "github.com/yincongcyincong/BasicStudy/library/const"
	"github.com/yincongcyincong/BasicStudy/model/dao"
	"github.com/yincongcyincong/BasicStudy/proto/api"
	"google.golang.org/protobuf/proto"
)

// AddStudyTest 新增信息

func AddStudyTest(c *gin.Context, input *api.AddStudyTestReq) (*api.AddStudyTestData, int, error) {
	// 参数校验

	if input == nil {
		logrus.Warningf("AddStudyTest input invalid, input=%+v", input)
		return nil, _const.ParamErrorNo, errors.New("input params invalid")
	}

	// 插入DB
	db := dao.NewStudyTestDao()
	insertReq := dao.StudyTestItem{}

	if input.Name != nil {
		insertReq.Name = proto.String(input.GetName())
	}

	id, err := db.Add(context.Background(), insertReq)
	if err != nil {
		logrus.Warningf("AddStudyTest insert db fail, err=%v, input=%+v", err, insertReq)
		return nil, _const.DBQueryError, err
	}

	data := &api.AddStudyTestData{
		Id: proto.Uint64(uint64(id)),
	}
	return data, _const.SuccNo, nil
}

// UpdateStudyTest 修改信息

func UpdateStudyTest(c *gin.Context, input *api.UpdateStudyTestReq) (int, error) {
	// 参数校验

	if input == nil || input.GetId() <= 0 {
		logrus.Warningf("UpdateStudyTest input invalid, input=%+v", input)
		return _const.ParamErrorNo, errors.New("input params invalid")
	}

	id := input.GetId()

	// 更新DB
	db := dao.NewStudyTestDao()
	updateReq := dao.StudyTestItem{}

	if input.Name != nil {
		updateReq.Name = proto.String(input.GetName())
	}

	_, err := db.Update(context.Background(), updateReq, id)
	if err != nil {
		logrus.Warningf("UpdateStudyTest update db fail, err=%v, input=%+v", err, updateReq)
		return _const.DBQueryError, err
	}
	return _const.SuccNo, nil
}

// DelStudyTest 删除信息

func DelStudyTest(c *gin.Context, input *api.DelStudyTestReq) (int, error) {
	// 参数校验

	if input == nil || input.GetId() <= 0 {
		logrus.Warningf("DelStudyTest input invalid, input=%+v", input)
		return _const.ParamErrorNo, errors.New("input params invalid")
	}

	id := input.GetId()

	// 更新DB
	db := dao.NewStudyTestDao()
	delReq := dao.StudyTestItem{}

	_, err := db.Update(context.Background(), delReq, id)
	if err != nil {
		logrus.Warningf("DelStudyTest update db fail, err=%v, input=%+v", err, delReq)
		return _const.DBQueryError, err
	}
	return _const.SuccNo, nil
}

// MgetStudyTestByIDs 按照ID批量查询信息

func MgetStudyTestByIDs(c *gin.Context, input *api.MgetStudyTestByIDsReq) (*api.MgetStudyTestByIDsData, int, error) {
	// 参数校验

	if input == nil || len(input.GetIds()) <= 0 {
		logrus.Warningf("MgetStudyTestByIDs input invalid, input=%+v", input)
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

	info := map[uint64]*api.StudyTest{}

	// 查询DB
	db := dao.NewStudyTestDao()
	queryRes, err := db.MgetByIDs(context.Background(), queryIDs)
	if err != nil {
		logrus.Warningf("MgetStudyTestByIDs query db fail, err=%v, query_id=%+v", err, &queryIDs)
		return nil, _const.DBQueryError, err
	}
	for _, res := range formatStudyTest(queryRes) {
		info[res.GetId()] = res
	}

	data := &api.MgetStudyTestByIDsData{
		Info: info,
	}
	return data, _const.SuccNo, nil
}

// MgetStudyTestByCond 按照条件查询信息（无缓存）

func MgetStudyTestByCond(c *gin.Context, input *api.MgetStudyTestByCondReq) (*api.MgetStudyTestByCondData, int, error) {
	// 参数校验
	if input == nil {
		logrus.Warningf("MgetStudyTestByCond input invalid, input=%+v", input)
		return nil, _const.ParamErrorNo, errors.New("input params invalid")
	}

	// 查询DB
	db := dao.NewStudyTestDao()
	queryReq := dao.StudyTestItem{}

	queryRes, totalNum, err := db.MgetByPage(context.Background(), queryReq, input.GetPn(), input.GetRn())
	if err != nil {
		logrus.Warningf("MgetStudyTestByCond query db fail, err=%v, input=%+v", err, input)
		return nil, _const.DBQueryError, err
	}

	data := &api.MgetStudyTestByCondData{
		Info:     formatStudyTest(queryRes),
		TotalNum: proto.Uint32(uint32(totalNum)),
	}
	return data, _const.SuccNo, nil
}

// formatStudyTest 格式化返回值，将内部使用的dao结构体转换为对外使用的proto结构体
func formatStudyTest(srcData []dao.StudyTestItem) []*api.StudyTest {
	dstData := make([]*api.StudyTest, 0, len(srcData))
	for _, data := range srcData {
		info := api.StudyTest{}
		if data.Id != nil {
			info.Id = proto.Uint64(*data.Id)
		}
		if data.Name != nil {
			info.Name = proto.String(*data.Name)
		}
		dstData = append(dstData, &info)
	}
	return dstData
}
