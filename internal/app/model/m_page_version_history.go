package model

import (
	"context"

	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
)

// IPageVersionHistory 页面历史版本存储接口
type IPageVersionHistory interface {
	// 查询数据
	Query(ctx context.Context, params schema.PageVersionHistoryQueryParam, opts ...schema.PageVersionHistoryQueryOptions) (*schema.PageVersionHistoryQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, id string, opts ...schema.PageVersionHistoryQueryOptions) (*schema.PageVersionHistory, error)
	// 创建数据
	Create(ctx context.Context, item schema.PageVersionHistory) error
	// 删除数据
	Delete(ctx context.Context, id string) error
}
