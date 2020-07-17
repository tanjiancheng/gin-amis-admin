package model

import (
	"context"

	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
)

// IPageManager 菜单管理存储接口
type IPageManager interface {
	// 查询数据
	Query(ctx context.Context, params schema.PageManagerQueryParam, opts ...schema.PageManagerQueryOptions) (*schema.PageManagerQueryResult, error)
	// 根据路由查询指定数据
	GetByRoute(ctx context.Context, id string, opts ...schema.PageManagerQueryOptions) (*schema.PageManager, error)
	// 查询指定数据
	Get(ctx context.Context, id string, opts ...schema.PageManagerQueryOptions) (*schema.PageManager, error)
	// 创建数据
	Create(ctx context.Context, item schema.PageManager) error
	// 更新数据
	Update(ctx context.Context, id string, item schema.PageManager) error
	// 删除数据
	Delete(ctx context.Context, id string) error
	// 更新状态
	UpdateStatus(ctx context.Context, id string, status int) error
	// 获取最后一条记录ID
	GetLastId(ctx context.Context) (int, error)
}
