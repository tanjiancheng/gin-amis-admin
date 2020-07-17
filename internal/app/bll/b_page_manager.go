package bll

import (
	"context"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
)

// IPageManager 页面管理业务逻辑接口
type IPageManager interface {
	// 初始化菜单数据
	InitData(ctx context.Context, dataFile string) error
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
	// 获取最后一条记录的ide
	GetLastId(ctx context.Context) (int, error)
	// 获取历史版本
	History(ctx context.Context, params schema.PageVersionHistoryQueryParam, opts ...schema.PageVersionHistoryQueryOptions) (*schema.PageVersionHistoryQueryResult, error)
	// 根据版本信息回滚
	Revert(ctx context.Context, pageVersionHistoryId string) error
}
