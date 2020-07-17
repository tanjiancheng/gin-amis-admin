package bll

import (
	"context"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
)

// IPageVersionHistory 页面历史管理业务逻辑接口
type IPageVersionHistory interface {
	// 查询数据
	Query(ctx context.Context, params schema.PageVersionHistoryQueryParam, opts ...schema.PageVersionHistoryQueryOptions) (*schema.PageVersionHistoryQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, id string, opts ...schema.PageVersionHistoryQueryOptions) (*schema.PageVersionHistory, error)
}
