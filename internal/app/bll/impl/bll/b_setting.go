package bll

import (
	"context"

	"github.com/tanjiancheng/gin-amis-admin/internal/app/bll"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/model"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/tanjiancheng/gin-amis-admin/pkg/errors"
	"github.com/google/wire"
)

var _ bll.ISetting = (*Setting)(nil)

// SettingSet 注入Setting
var SettingSet = wire.NewSet(wire.Struct(new(Setting), "*"), wire.Bind(new(bll.ISetting), new(*Setting)))

// Setting 示例程序
type Setting struct {
	TransModel   model.ITrans
	SettingModel model.ISetting
}

// Query 查询数据
func (a *Setting) Query(ctx context.Context) (*schema.SettingQueryResult, error) {
	return a.SettingModel.Query(ctx)
}

// Get 查询指定数据
func (a *Setting) Get(ctx context.Context, id string) (*schema.Setting, error) {
	item, err := a.SettingModel.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Create 创建数据
func (a *Setting) Create(ctx context.Context, item schema.Setting) error {
	err := a.SettingModel.Create(ctx, item)
	if err != nil {
		return err
	}
	return nil
}

// Update 更新数据
func (a *Setting) Update(ctx context.Context, id string, item schema.Setting) error {
	oldItem, err := a.SettingModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	item.CreatedAt = oldItem.CreatedAt

	return a.SettingModel.Update(ctx, id, item)
}

// Delete 删除数据
func (a *Setting) Delete(ctx context.Context, id string) error {
	oldItem, err := a.SettingModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.SettingModel.Delete(ctx, id)
}

func (a *Setting) Truncate(ctx context.Context) error {
	return a.SettingModel.Truncate(ctx)
}
