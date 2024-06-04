package dao

// study_question表 相关dao
// Code generated by dev-tools-cli.

import (
	"context"
	"time"

	"google.golang.org/protobuf/proto"
)

// StudyQuestionDao 表结构体
type StudyQuestionDao struct {
	DBDaoBase
}

// TableStudyQuestion 表名称
const TableStudyQuestion = "study_question"

const (
	// studyQuestionDefaultPn 分页查询使用的默认页码
	studyQuestionDefaultPn = 1

	// studyQuestionDefaultRn 分页查询使用的默认数据量
	studyQuestionDefaultRn = 20
)

// studyQuestionSelectFields 表中字段
var studyQuestionSelectFields = []string{
	"id", "create_time", "question", "answer", "test_id",
	"type",
}

// NewStudyQuestionDao 初始化
func NewStudyQuestionDao() *StudyQuestionDao {
	return &StudyQuestionDao{
		DBDaoBase: NewDBDao(TableStudyQuestion),
	}
}

// StudyQuestionItem is a mapping object for basic_study.study_question table
type StudyQuestionItem struct {
	Id         *uint64 `ddb:"id" json:"id"`                   // id
	CreateTime *uint64 `ddb:"create_time" json:"create_time"` // create time
	Question   *string  `ddb:"question" json:"question"`       // question
	Answer     *string  `ddb:"answer" json:"answer"`           // answer
	TestId     *uint64 `ddb:"test_id" json:"test_id"`         // test
	Type       *uint8  `ddb:"type" json:"type"`               // 1-math 2-english
}

// Add 新增信息
func (dao *StudyQuestionDao) Add(ctx context.Context, data StudyQuestionItem) (int64, error) {
	tm := time.Now().Unix()
	data.CreateTime = proto.Uint64(uint64(tm))

	lastInsertID, err := dao.BaseInsert(ctx, data)
	if err == nil && lastInsertID != 0 {
		return lastInsertID, nil
	}
	return 0, err
}

// Update 修改信息
func (dao *StudyQuestionDao) Update(ctx context.Context, data StudyQuestionItem, id uint64) (int64, error) {
	where := map[string]any{
		"id": id,
	}
	return dao.BaseUpdate(ctx, where, data)
}

// MgetByIDs 按照ID批量查询信息
func (dao *StudyQuestionDao) MgetByIDs(ctx context.Context, ids []uint64) ([]StudyQuestionItem, error) {
	var res []StudyQuestionItem
	where := map[string]any{
		"id in": ids,
	}

	selectFields := studyQuestionSelectFields

	err := dao.BaseSelectConvert(ctx, where, selectFields, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MgetByPage 按照指定条件查询信息
func (dao *StudyQuestionDao) MgetByPage(ctx context.Context, data StudyQuestionItem, pn, rn uint32) ([]StudyQuestionItem, int64, error) {
	where := map[string]any{}


	totalNum, err := dao.Count(ctx, where)
	if err != nil {
		return nil, 0, err
	}

	var res []StudyQuestionItem
	queryPn := uint(studyQuestionDefaultPn)
	if pn > 0 {
		queryPn = uint(pn)
	}
	queryRn := uint(studyQuestionDefaultRn)
	if rn > 0 {
		queryRn = uint(rn)
	}
	where["_limit"] = []uint{(queryPn - 1) * queryRn, queryRn}
	err = dao.BaseSelectConvert(ctx, where, studyQuestionSelectFields, &res)
	if err != nil {
		return nil, 0, err
	}
	return res, totalNum, nil
}
