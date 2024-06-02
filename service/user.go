package service

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	_const "github.com/yincongcyincong/BasicStudy/library/const"
	"github.com/yincongcyincong/BasicStudy/model/dao"
	"github.com/yincongcyincong/BasicStudy/proto/api"
	"google.golang.org/protobuf/proto"
)

// AddStudyUser 新增信息

func AddStudyUser(ctx context.Context, input *api.AddStudyUserReq) (*api.AddStudyUserData, int, error) {
	// 参数校验

	if input == nil {
		logrus.Warningf("AddStudyUser input invalid, input=%+v", input)
		return nil, _const.ParamErrorNo, errors.New("input params invalid")
	}

	// 插入DB
	db := dao.NewStudyUserDao()
	insertReq := dao.StudyUserItem{}
	if input.Username != nil {
		insertReq.Username = proto.String(input.GetUsername())
	}
	if input.Password != nil {
		insertReq.Password = proto.String(input.GetPassword())
	}
	if input.Status != nil {
		infoStatus := uint8(input.GetStatus())
		insertReq.Status = &infoStatus
	}

	id, err := db.Add(ctx, insertReq)
	if err != nil {
		logrus.Warningf("AddStudyUser insert db fail, err=%v, input=%+v", err, insertReq)
		return nil, _const.DBQueryError, err
	}

	data := &api.AddStudyUserData{
		Id: proto.Uint64(uint64(id)),
	}
	return data, _const.SuccNo, nil
}

// UpdateStudyUser 修改信息

func UpdateStudyUser(ctx context.Context, input *api.UpdateStudyUserReq) (int, error) {
	// 参数校验

	if input == nil || input.GetId() <= 0 {
		logrus.Warningf("UpdateStudyUser input invalid, input=%+v", input)
		return _const.ParamErrorNo, errors.New("input params invalid")
	}

	id := input.GetId()

	// 更新DB
	db := dao.NewStudyUserDao()
	updateReq := dao.StudyUserItem{}
	if input.Password != nil {
		updateReq.Password = proto.String(input.GetPassword())
	}
	if input.Status != nil {
		infoStatus := uint8(input.GetStatus())
		updateReq.Status = &infoStatus
	}

	_, err := db.Update(ctx, updateReq, id)
	if err != nil {
		logrus.Warningf("UpdateStudyUser update db fail, err=%v, input=%+v", err, updateReq)
		return _const.DBQueryError, err
	}
	return _const.SuccNo, nil
}

// DelStudyUser 删除信息

func DelStudyUser(ctx context.Context, input *api.DelStudyUserReq) (int, error) {
	// 参数校验
	if input == nil || input.GetId() <= 0 {
		logrus.Warningf("DelStudyUser input invalid, input=%+v", input)
		return _const.ParamErrorNo, errors.New("input params invalid")
	}

	id := input.GetId()

	// 更新DB
	db := dao.NewStudyUserDao()
	delReq := dao.StudyUserItem{}

	_, err := db.Update(ctx, delReq, id)
	if err != nil {
		logrus.Warningf("DelStudyUser update db fail, err=%v, input=%+v", err, delReq)
		return _const.DBQueryError, err
	}
	return _const.SuccNo, nil
}

// MgetStudyUserByIDs 按照ID批量查询信息

func MgetStudyUserByIDs(ctx context.Context, input *api.MgetStudyUserByIDsReq) (*api.MgetStudyUserByIDsData, int, error) {
	// 参数校验

	if input == nil || len(input.GetIds()) <= 0 {
		logrus.Warningf("MgetStudyUserByIDs input invalid, input=%+v", input)
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

	info := map[uint64]*api.StudyUser{}

	// 查询DB
	db := dao.NewStudyUserDao()
	queryRes, err := db.MgetByIDs(ctx, queryIDs)
	if err != nil {
		logrus.Warningf("MgetStudyUserByIDs query db fail, err=%v, query_id=%+v", err, queryIDs)
		return nil, _const.DBQueryError, err
	}
	for _, res := range formatStudyUser(queryRes) {
		info[res.GetId()] = res
	}

	data := &api.MgetStudyUserByIDsData{
		Info: info,
	}
	return data, _const.SuccNo, nil
}

// MgetStudyUserByCond 按照条件查询信息（无缓存）

func MgetStudyUserByCond(ctx context.Context, input *api.MgetStudyUserByCondReq) (*api.MgetStudyUserByCondData, int, error) {
	// 参数校验

	if input == nil {
		logrus.Warningf("MgetStudyUserByCond input invalid, input=%+v", input)
		return nil, _const.ParamErrorNo, errors.New("input params invalid")
	}

	// 查询DB
	db := dao.NewStudyUserDao()
	queryReq := dao.StudyUserItem{}

	queryRes, totalNum, err := db.MgetByPage(ctx, queryReq, input.GetPn(), input.GetRn())
	if err != nil {
		logrus.Warningf("MgetStudyUserByCond query db fail, err=%v, input=%+v", err, input)
		return nil, _const.DBQueryError, err
	}

	data := &api.MgetStudyUserByCondData{
		Info:     formatStudyUser(queryRes),
		TotalNum: proto.Uint32(uint32(totalNum)),
	}
	return data, _const.SuccNo, nil
}

// formatStudyUser 格式化返回值，将内部使用的dao结构体转换为对外使用的proto结构体
func formatStudyUser(srcData []dao.StudyUserItem) []*api.StudyUser {
	dstData := make([]*api.StudyUser, 0, len(srcData))
	for _, data := range srcData {
		info := api.StudyUser{}
		if data.Id != nil {
			info.Id = proto.Uint64(*data.Id)
		}
		if data.CreateTime != nil {
			info.CreateTime = proto.Uint64(*data.CreateTime)
		}
		if data.Username != nil {
			info.Username = proto.String(*data.Username)
		}
		if data.Password != nil {
			info.Password = proto.String(*data.Password)
		}
		if data.Status != nil {
			infoStatus := uint32(*data.Status)
			info.Status = proto.Uint32(infoStatus)
		}
		dstData = append(dstData, &info)
	}
	return dstData
}
