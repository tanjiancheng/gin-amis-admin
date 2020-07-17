package entity

import (
	"context"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/tanjiancheng/gin-amis-admin/pkg/util"
	"github.com/jinzhu/gorm"
)

// GetPageVersionHistory 获取页面版本存储
func GetPageVersionHistoryDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(PageVersionHistory))
}

// SchemaPageVersionHistory 页面版本对象
type SchemaPageVersionHistory schema.PageVersionHistory

// ToPageVersionHistory 转换为页面版本实体
func (a SchemaPageVersionHistory) ToPageVersionHistory() *PageVersionHistory {
	item := new(PageVersionHistory)
	util.StructMapToStruct(a, item)
	return item
}

// PageVersionHistory 页面版本实体
type PageVersionHistory struct {
	ID              int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`            // 自增ID
	PageManagerId   int    `gorm:"column:page_manager_id;default:0;not null;"`      // 页面标识
	PageManagerInfo string `gorm:"column:page_manger_info;type:longtext;not null;"` // 页面源码
	CreateTime      int64  `gorm:"column:create_time;"`                             // 创建时间
}

// TableName 表名
func (a PageVersionHistory) TableName() string {
	return Model{}.TableName("page_version_history")
}

// ToSchemaPageVersionHistory 转换为页面版本对象
func (a PageVersionHistory) ToSchemaPageVersionHistory() *schema.PageVersionHistory {
	item := new(schema.PageVersionHistory)
	util.StructMapToStruct(a, item)
	return item
}

type PageVersionHistorys []*PageVersionHistory

// ToSchemaPageVersionHistorys 转换为页面版本对象列表
func (a PageVersionHistorys) ToSchemaPageVersionHistorys() []*schema.PageVersionHistory {
	list := make([]*schema.PageVersionHistory, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaPageVersionHistory()
	}
	return list
}
