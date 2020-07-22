package entity

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/config"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/tanjiancheng/gin-amis-admin/pkg/util"
)

type GPlatform struct {
	ID         int    `gorm:"column:id;primary_key;auto_increment;"`
	AppID      string `gorm:"column:app_id;unique_index;size:36;not null;"`    // 平台ID
	Name       string `gorm:"column:name;size:100;index;default:'';not null;"` // 平台名称
	Status     int    `gorm:"column:status;default:0;not null;"`               // 状态(1:启用 -1:停用)
	CreateTime int64  `gorm:"column:create_time;"`                             // 创建时间
}

func (a GPlatform) ToSchemaGPlatform() *schema.GPlatform {
	item := new(schema.GPlatform)
	util.StructMapToStruct(a, item)
	return item
}

// TableName 表名
func (a GPlatform) TableName() string {
	return fmt.Sprintf("%s%s", config.C.Gorm.GlobalTablePrefix, "platform")
}

type SchemaGPlatform schema.GPlatform

func (a SchemaGPlatform) ToGPlatform() *GPlatform {
	item := new(GPlatform)
	util.StructMapToStruct(a, item)
	return item
}

func GetGPlatformDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(GPlatform))
}

type GPlatforms []*GPlatform

func (a GPlatforms) ToSchemaGPlatforms() []*schema.GPlatform {
	list := make([]*schema.GPlatform, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaGPlatform()
	}
	return list
}
