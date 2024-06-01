package dao

import (
	"context"
	"database/sql"
	"github.com/yincongcyincong/BasicStudy/bootstrap"
)

// DBDaoBase 基础dao结构体
type DBDaoBase struct {
	TableName string
	DB        *sql.DB
}

// NewDBDao 新建 Dao
func NewDBDao(tableName string) DBDaoBase {
	return DBDaoBase{
		TableName: tableName,
	}
}

// BaseSelectConvert 查询并将数据反射赋值到结构体
func (dao *DBDaoBase) BaseSelectConvert(ctx context.Context, where map[string]any, fields []string, target any) error {
	_, err := bootstrap.DB.Table(dao.TableName).Where(where).Cols(fields...).Get(target)
	return err
}

// BaseUpdate 更新数据
func (dao *DBDaoBase) BaseUpdate(ctx context.Context, where map[string]any, assigns any) (int64, error) {
	return bootstrap.DB.Table(dao.TableName).Where(where).Update(assigns)
}

// BaseInsert 插入数据
func (dao *DBDaoBase) BaseInsert(ctx context.Context, data any) (lastInsertID int64, err error) {
	return bootstrap.DB.Table(dao.TableName).Insert(data)
}

// BaseDelete 删除数据
func (dao *DBDaoBase) BaseDelete(ctx context.Context, where map[string]any) (int64, error) {
	return bootstrap.DB.Table(dao.TableName).Where(where).Delete()
}

// Count 求总数
func (dao *DBDaoBase) Count(ctx context.Context, where map[string]any) (int64, error) {
	return bootstrap.DB.Table(dao.TableName).Where(where).Count()
}

// BaseSQLQuery 自定义sql query
func (dao *DBDaoBase) BaseSQLQuery(ctx context.Context, sql string, data map[string]any, target any) error {
	_, err := bootstrap.DB.SQL(sql).Get(target)
	return err
}

// BaseSQLExec 自定义sql exec
func (dao *DBDaoBase) BaseSQLExec(ctx context.Context, sql string) (affected int64, lastInsertID int64, err error) {
	res, err := bootstrap.DB.Exec(sql)
	if err != nil {
		return 0, 0, err
	}
	affectedRow, _ := res.RowsAffected()
	lastId, _ := res.LastInsertId()
	return affectedRow, lastId, err
}
