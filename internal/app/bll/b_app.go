package bll

import (
	"context"
)

type IApp interface {
	// 初始化数据
	Init(ctx context.Context, appId string) error
	// 查询是否初始化过
	Query(ctx context.Context, appId string) (bool, error)
}
