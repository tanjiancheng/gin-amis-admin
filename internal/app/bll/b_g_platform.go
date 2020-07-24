package bll

import (
	"context"

	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
)

// IGPlatform 业务逻辑接口
type IGPlatform interface {
	// 查询数据
	Query(ctx context.Context, params schema.GPlatformQueryParam, opts ...schema.GPlatformQueryOptions) (*schema.GPlatformQueryResult, error)
	// 检查appId是否合法
	Check(ctx context.Context, appId string) error
	// 查询指定数据
	Get(ctx context.Context, id string, opts ...schema.GPlatformQueryOptions) (*schema.GPlatform, error)
	// 获取select的查询配置
	GetOptions(ctx context.Context) (*schema.GPlatformSelectOptions, error)
	// 根据app_id查询对应的数据
	GetByAppId(ctx context.Context, appId string) (*schema.GPlatform, error)
	// 创建数据
	Create(ctx context.Context, item schema.GPlatform)  error
	// 更新数据
	Update(ctx context.Context, id string, item schema.GPlatform) error
	// 删除数据
	Delete(ctx context.Context, id string) error
	// 更新状态
	UpdateStatus(ctx context.Context, id string, status int) error
}
