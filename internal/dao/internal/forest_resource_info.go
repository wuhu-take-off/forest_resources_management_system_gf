// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ForestResourceInfoDao is the data access object for table forest_resource_info.
type ForestResourceInfoDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ForestResourceInfoColumns // columns contains all the column names of Table for convenient usage.
}

// ForestResourceInfoColumns defines and stores column names for table forest_resource_info.
type ForestResourceInfoColumns struct {
	ForestResourceInfoId      string // 林地资源信息主键
	ForestId                  string // 林地id
	TreeSpeciesId             string // 树种id
	ForestResourceExamineTime string // 检查日期
	ForestResourceExamineType string // 检查类型
	ForestStatus              string // 林地状态
	ForestResourceExamineInfo string // 检测数据
	ProblemDes                string // 问题描述
	TreatmentSuggestion       string // 处理建议
	ScenePhoto                string // 现场照片
	SceneVideo                string // 视频记录
}

// forestResourceInfoColumns holds the columns for table forest_resource_info.
var forestResourceInfoColumns = ForestResourceInfoColumns{
	ForestResourceInfoId:      "forest_resource_info_id",
	ForestId:                  "forest_id",
	TreeSpeciesId:             "tree_species_id",
	ForestResourceExamineTime: "forest_resource_examine_time",
	ForestResourceExamineType: "forest_resource_examine_type",
	ForestStatus:              "forest_status",
	ForestResourceExamineInfo: "forest_resource_examine_info",
	ProblemDes:                "problem_des",
	TreatmentSuggestion:       "treatment_suggestion",
	ScenePhoto:                "scene_photo",
	SceneVideo:                "scene_video",
}

// NewForestResourceInfoDao creates and returns a new DAO object for table data access.
func NewForestResourceInfoDao() *ForestResourceInfoDao {
	return &ForestResourceInfoDao{
		group:   "default",
		table:   "forest_resource_info",
		columns: forestResourceInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ForestResourceInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ForestResourceInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ForestResourceInfoDao) Columns() ForestResourceInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ForestResourceInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ForestResourceInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ForestResourceInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
