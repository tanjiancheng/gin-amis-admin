package schema

import (
	"github.com/tanjiancheng/gin-amis-admin/pkg/util"
)

// PageVersionHistory 菜单对象
type PageVersionHistory struct {
	ID              int    `json:"id"`               // 唯一标识
	Name            string `json:"name"`                // 页面名称
	PageManagerId   int    `json:"page_manager_id"`  // 页面标识
	PageManagerInfo string `json:"page_manger_info"` // 页面历史信息
	CreateTime      int64  `json:"history_create_time"`      // 页面版本创建时间
	PageCreateTime  int64  `json:"create_time"`                // 页面创建时间
}

type PageVersionHistorys []*PageVersionHistory

func (a *PageVersionHistory) String() string {
	return util.JSONMarshalToString(a)
}

// PageVersionHistoryQueryParam 查询条件
type PageVersionHistoryQueryParam struct {
	PaginationParam
	PageManagerId string `form:"page_manager_id"` // 页面id
}

// PageVersionHistoryQueryOptions 查询可选参数项
type PageVersionHistoryQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// PageVersionHistoryQueryResult 查询结果
type PageVersionHistoryQueryResult struct {
	Data       PageVersionHistorys
	PageResult *PaginationResult
}
