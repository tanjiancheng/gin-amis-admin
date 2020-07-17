package bll

import (
	"context"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/bll"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/model"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/tanjiancheng/gin-amis-admin/pkg/errors"
	"github.com/google/wire"
)

var _ bll.IPageVersionHistory = (*PageVersionHistory)(nil)

// PageVersionHistorySet 注入PageVersionHistory
var PageVersionHistorySet = wire.NewSet(wire.Struct(new(PageVersionHistory), "*"), wire.Bind(new(bll.IPageVersionHistory), new(*PageVersionHistory)))

// 页面管理
type PageVersionHistory struct {
	TransModel              model.ITrans
	PageVersionHistoryModel model.IPageVersionHistory
}

// 查询数据
func (a *PageVersionHistory) Query(ctx context.Context, params schema.PageVersionHistoryQueryParam, opts ...schema.PageVersionHistoryQueryOptions) (*schema.PageVersionHistoryQueryResult, error) {
	return a.PageVersionHistoryModel.Query(ctx, params, opts...)
}

// 查询指定数据
func (a *PageVersionHistory) Get(ctx context.Context, id string, opts ...schema.PageVersionHistoryQueryOptions) (*schema.PageVersionHistory, error) {
	item, err := a.PageVersionHistoryModel.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}
