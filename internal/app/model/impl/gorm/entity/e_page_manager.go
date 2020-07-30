package entity

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/iutil"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/tanjiancheng/gin-amis-admin/pkg/util"
)

// GetPageManagerDb 获取页面管理存储
func GetPageManagerDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(PageManager))
}

// SchemaPageManager 页面管理对象
type SchemaPageManager schema.PageManager

// ToPageManager 转换为页面管理实体
func (a SchemaPageManager) ToPageManager() *PageManager {
	item := new(PageManager)
	util.StructMapToStruct(a, item)
	return item
}

// PageManager 页面管理实体
type PageManager struct {
	ID         int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`                      // 自增ID
	Identify   string `gorm:"unique_index;column:identify;size:64;default:'';not null;"` // 页面标识
	Name       string `gorm:"column:name;size:64;default:'';not null;"`                  // 页面名称
	Meta       string `gorm:"column:meta;type:longtext;not null;"`                       // 页面元信息
	Source     string `gorm:"column:source;type:longtext;not null;"`                     // 页面源码
	Creator    string `gorm:"column:creator;size:32;default:'';not null;"`               // 创建者
	CreateTime int64  `gorm:"column:create_time;"`                                       // 创建时间
	ModifyTime int64  `gorm:"column:modify_time;"`                                       // 更新时间
}

// TableName 表名
func (a PageManager) TableName() string {
	return Model{}.TableName("page_manager")
}

// ToSchemaPageManager 转换为页面管理对象
func (a PageManager) ToSchemaPageManager() *schema.PageManager {
	meta := a.Meta
	source := a.Source
	jsonPathx := new(iutil.JsonPathx)
	item := new(schema.PageManager)
	item.RenderSource, _ = jsonPathx.Search(meta, source)
	util.StructMapToStruct(a, item)
	return item
}

type PageManagers []*PageManager

// ToSchemaPageManagers 转换为页面管理对象列表
func (a PageManagers) ToSchemaPageManagers() []*schema.PageManager {
	list := make([]*schema.PageManager, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaPageManager()
	}
	return list
}
