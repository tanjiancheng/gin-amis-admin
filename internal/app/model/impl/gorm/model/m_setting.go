package model

import (
	"context"

	"github.com/tanjiancheng/gin-amis-admin/internal/app/model"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/model/impl/gorm/entity"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/tanjiancheng/gin-amis-admin/pkg/errors"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var _ model.ISetting = (*Setting)(nil)

// SettingSet 注入Setting
var SettingSet = wire.NewSet(wire.Struct(new(Setting), "*"), wire.Bind(new(model.ISetting), new(*Setting)))

// Setting 示例存储
type Setting struct {
	DB *gorm.DB
}

// Query 查询数据
func (a *Setting) Query(ctx context.Context) (*schema.SettingQueryResult, error) {
	db := entity.GetSettingDB(ctx, a.DB)
	var list entity.Settings
	err := db.Find(&list).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.SettingQueryResult{
		Data: list.ToSchemaSettings(),
	}
	return qr, nil
}

// Get 查询指定数据
func (a *Setting) Get(ctx context.Context, id string) (*schema.Setting, error) {
	db := entity.GetSettingDB(ctx, a.DB).Where("key=?", id)
	var item entity.Setting
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaSetting(), nil
}

// Create 创建数据
func (a *Setting) Create(ctx context.Context, item schema.Setting) error {
	eitem := entity.SchemaSetting(item).ToSetting()
	db := entity.GetSettingDB(ctx, a.DB)
	result := db.Create(eitem)
	return errors.WithStack(result.Error)
}

// Update 更新数据
func (a *Setting) Update(ctx context.Context, id string, item schema.Setting) error {
	eitem := entity.SchemaSetting(item).ToSetting()
	result := entity.GetSettingDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

// Delete 删除数据
func (a *Setting) Delete(ctx context.Context, id string) error {
	result := entity.GetSettingDB(ctx, a.DB).Where("id=?", id).Delete(entity.Setting{})
	return errors.WithStack(result.Error)
}

func (a *Setting) Truncate(ctx context.Context) error {
	db := entity.GetSettingDB(ctx, a.DB).Exec("truncate table " + (entity.Setting{}).TableName())
	if db.Error != nil {
		return db.Error
	}
	return nil
}
