package model

import (
	"context"

	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
)

type IGTplMall interface {
	// 查询数据
	Query(ctx context.Context, params schema.GTplMallQueryParam, opts ...schema.GTplMallQueryOptions) (*schema.GTplMallQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, id string, opts ...schema.GTplMallQueryOptions) (*schema.GTplMall, error)
	// 根据appId查询数据
	GetByIdentify(ctx context.Context, identify string) (*schema.GTplMall, error)
	// 创建数据
	Create(ctx context.Context, item schema.GTplMall) error
	// 更新数据
	Update(ctx context.Context, id string, item schema.GTplMall) error
	// 删除数据
	Delete(ctx context.Context, id string) error
	// 更新状态
	UpdateStatus(ctx context.Context, id string, status int) error
}
