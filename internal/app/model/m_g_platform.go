package model

import (
	"context"

	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
)

type IGPlatform interface {
	// 查询数据
	Query(ctx context.Context, params schema.GPlatformQueryParam, opts ...schema.GPlatformQueryOptions) (*schema.GPlatformQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, id string, opts ...schema.GPlatformQueryOptions) (*schema.GPlatform, error)
	// 查询select的选项
	GetOptions(ctx context.Context) (*schema.GPlatformSelectOptions, error)
	// 根据appId查询数据
	GetByAppId(ctx context.Context, appId string) (*schema.GPlatform, error)
	// 创建数据
	Create(ctx context.Context, item schema.GPlatform) error
	// 更新数据
	Update(ctx context.Context, id string, item schema.GPlatform) error
	// 删除数据
	Delete(ctx context.Context, id string) error
	// 更新状态
	UpdateStatus(ctx context.Context, id string, status int) error
}
