package bll

import (
	"context"
	"regexp"

	"github.com/google/wire"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/bll"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/model"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/tanjiancheng/gin-amis-admin/pkg/errors"
)

var _ bll.IGPlatform = (*GPlatform)(nil)

// GPlatformSet 注入GPlatform
var GPlatformSet = wire.NewSet(wire.Struct(new(GPlatform), "*"), wire.Bind(new(bll.IGPlatform), new(*GPlatform)))

// GPlatform 示例程序
type GPlatform struct {
	GPlatformModel model.IGPlatform
}

// Query 查询数据
func (a *GPlatform) Query(ctx context.Context, params schema.GPlatformQueryParam, opts ...schema.GPlatformQueryOptions) (*schema.GPlatformQueryResult, error) {
	return a.GPlatformModel.Query(ctx, params, opts...)
}

func (a *GPlatform) Check(ctx context.Context, appId string) error {
	if len(appId) <= 0 {
		return errors.NewResponse(-1, 200, "appId不合法")
	}
	r, _ := regexp.Compile("[A-Za-z0-9]")
	if !r.MatchString(appId) {
		return errors.NewResponse(-1, 200, "appId必须是[A-Za-z0-9]的格式")
	}
	//判断是否是数字或者字母
	return nil
}

// Get 查询指定数据
func (a *GPlatform) Get(ctx context.Context, id string, opts ...schema.GPlatformQueryOptions) (*schema.GPlatform, error) {
	item, err := a.GPlatformModel.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *GPlatform) GetByAppId(ctx context.Context, appId string) (*schema.GPlatform, error) {
	item, err := a.GPlatformModel.GetByAppId(ctx, appId)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Create 创建数据
func (a *GPlatform) Create(ctx context.Context, item schema.GPlatform) error {
	//判断是否已经插入
	oldItem, err := a.GPlatformModel.GetByAppId(ctx, item.AppID)
	if oldItem != nil && oldItem.ID > 0 { //创建过则不做任何处理
		return nil
	}
	err = a.GPlatformModel.Create(ctx, item)
	if err != nil {
		return err
	}
	return nil
}

// Update 更新数据
func (a *GPlatform) Update(ctx context.Context, id string, item schema.GPlatform) error {
	oldItem, err := a.GPlatformModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	item.ID = oldItem.ID

	return a.GPlatformModel.Update(ctx, id, item)
}

// Delete 删除数据
func (a *GPlatform) Delete(ctx context.Context, id string) error {
	oldItem, err := a.GPlatformModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.GPlatformModel.Delete(ctx, id)
}

// UpdateStatus 更新状态
func (a *GPlatform) UpdateStatus(ctx context.Context, id string, status int) error {
	oldItem, err := a.GPlatformModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.GPlatformModel.UpdateStatus(ctx, id, status)
}
