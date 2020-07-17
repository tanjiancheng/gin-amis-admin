package bll

import (
	"context"

	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
)

// ISetting demo业务逻辑接口
type ISetting interface {
	// 查询数据
	Query(ctx context.Context) (*schema.SettingQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, id string) (*schema.Setting, error)
	// 创建数据
	Create(ctx context.Context, item schema.Setting) error
	// 更新数据
	Update(ctx context.Context, id string, item schema.Setting) error
	// 删除数据
	Delete(ctx context.Context, id string) error
	// 清空表
	Truncate(ctx context.Context) error
}
