package schema

import (
	"github.com/tanjiancheng/gin-amis-admin/pkg/util"
)

// PageManager 菜单对象
type PageManager struct {
	ID           int    `yaml:"-" json:"id"`              // 唯一标识
	Identify     string `yaml:"identify" json:"identify"` // 页面标识
	Name         string `yaml:"name" json:"name"`         // 页面名称
	Meta         string `yaml:"meta" json:"meta"`         // 页面元信息
	Source       string `yaml:"source" json:"source"`     // 页面源码
	//OriginSource string `yaml:"-" json:"origin_source"`   // 原始页面源码
	RenderSource string `yaml:"-" json:"render_source"`   // 编译后页面源码
	Creator      string `yaml:"creator" json:"creator"`   // 创建者
	CreateTime   int64  `yaml:"-" json:"create_time"`     // 创建时间
	UpdateTime   int64  `yaml:"-" json:"modify_time"`     // 更新时间
}

type PageManagers []*PageManager

func (a *PageManager) String() string {
	return util.JSONMarshalToString(a)
}

// PageManagerQueryParam 查询条件
type PageManagerQueryParam struct {
	PaginationParam
	Route string `form:"route"`      // 页面标识
	Name  string `form:"queryValue"` // 页面名称
}

// PageManagerQueryOptions 查询可选参数项
type PageManagerQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// PageManagerQueryResult 查询结果
type PageManagerQueryResult struct {
	Data       PageManagers
	PageResult *PaginationResult
}
