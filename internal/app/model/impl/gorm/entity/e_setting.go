package entity

import (
	"context"
	"time"

	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/tanjiancheng/gin-amis-admin/pkg/util"
	"github.com/jinzhu/gorm"
)

// GetSettingDB 获取demo存储
func GetSettingDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Setting))
}

// SchemaSetting demo对象
type SchemaSetting schema.Setting

// ToSetting 转换为demo实体
func (a SchemaSetting) ToSetting() *Setting {
	item := new(Setting)
	util.StructMapToStruct(a, item)
	return item
}

// Setting demo实体
type Setting struct {
	Key       string    `gorm:"column:key;size:64;unique;default:'';not null;"` // 键
	Value     string    `gorm:"column:value;type:longtext;not null;"`           // 值
	CreatedAt time.Time `gorm:"column:created_at;index;"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;"`
}

// TableName 表名
func (a Setting) TableName() string {
	return Model{}.TableName("setting")
}

// ToSchemaSetting 转换为demo对象
func (a Setting) ToSchemaSetting() *schema.Setting {
	item := new(schema.Setting)
	util.StructMapToStruct(a, item)
	return item
}

// Settings demo列表
type Settings []*Setting

// ToSchemaSettings 转换为demo对象列表
func (a Settings) ToSchemaSettings() []*schema.Setting {
	list := make([]*schema.Setting, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaSetting()
	}
	return list
}
