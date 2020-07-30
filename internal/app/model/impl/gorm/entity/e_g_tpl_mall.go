package entity

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/config"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/tanjiancheng/gin-amis-admin/pkg/util"
)

type GTplMall struct {
	ID         int    `gorm:"column:id;primary_key;auto_increment;"`                     // 自增ID
	Identify   string `gorm:"unique_index;column:identify;size:64;default:'';not null;"` // 模板标识
	Scope      string `gorm:"column:scope;size:256;index;not null;"`                     // 应用限制 *为所有使用，其他情况为具体的app_id下才能可见
	Name       string `gorm:"column:name;size:128;index;default:'';not null;"`           // 模板名称
	Desc       string `gorm:"column:desc;size:256;default:'';not null;"`                 // 模板说明
	Meta       string `gorm:"column:meta;type:longtext;not null;"`                       // 页面元信息
	Source     string `gorm:"column:source;type:longtext;not null;"`                     // 页面源码
	MockData   string `gorm:"column:mock_data;type:longtext;not null;"`                  // mock接口数据
	Icon       string `gorm:"column:icon;size:32;default:'';not null;"`                  // 模板图标
	Status     int    `gorm:"column:status;default:0;not null;"`                         // 模板状态 0未发布 1正常 -1禁用
	Creator    string `gorm:"column:creator;size:32;default:'';not null;"`               // 创建者
	CreateTime int64  `gorm:"column:create_time;"`                                       // 创建时间
	UpdateTime int64  `gorm:"column:update_time;"`                                       // 更新时间
}

func (a GTplMall) ToSchemaGTplMall() *schema.GTplMall {
	mockData := a.MockData
	item := new(schema.GTplMall)
	_ = json.Unmarshal([]byte(mockData), &item.MockData)
	util.StructMapToStruct(a, item)
	return item
}

// TableName 表名
func (a GTplMall) TableName() string {
	return fmt.Sprintf("%s%s", config.C.Gorm.GlobalTablePrefix, "tpl_mall")
}

type SchemaGTplMall schema.GTplMall

func (a SchemaGTplMall) ToGTplMall() *GTplMall {
	mockData := a.MockData
	mockDataJsonStr, _ := json.Marshal(mockData)
	item := new(GTplMall)
	item.MockData = string(mockDataJsonStr)
	util.StructMapToStruct(a, item)
	return item
}

func GetGTplMallDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(GTplMall))
}

type GTplMalls []*GTplMall

func (a GTplMalls) ToSchemaGTplMalls() []*schema.GTplMall {
	list := make([]*schema.GTplMall, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaGTplMall()
	}
	return list
}
