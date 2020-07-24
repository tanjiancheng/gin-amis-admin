package schema

type GTplMall struct {
	ID         int            `yaml:"-" json:"id"`                // 自增ID
	Identify   string         `yaml:"identify" json:"identify"`   // 模板标识
	Scope      string         `yaml:"scope" json:"scope"`         // 应用限制 *为所有使用，其他情况为具体的app_id下才能可见
	Name       string         `yaml:"name" json:"name"`           // 模板名abc称
	Desc       string         `yaml:"desc" json:"desc"`           // 模板说明
	Source     string         `yaml:"source" json:"source"`       // 页面源码
	MockData   []MockDataItem `yaml:"mock_data" json:"mock_data"` // mock接口数据
	Icon       string         `yaml:"icon" json:"icon"`           // 模板图标
	IconFull   string         `yaml:"-" json:"icon_full"`         // 带i标签的模板图标
	Status     int            `yaml:"status" json:"status"`       //模板状态
	Creator    string         `yaml:"creator" json:"creator"`     // 创建者
	CreateTime int64          `yaml:"-" json:"create_time"`       // 创建时间
	UpdateTime int64          `yaml:"-" json:"update_time"`       // 更新时间
}

type GTplMalls []*GTplMall

type MockDataItem struct {
	Path string `json:"path"`
	Data string `json:"data"`
}

type GTplMallQueryParam struct {
	PaginationParam
	QueryValue string `form:"queryValue"` // 查询值
	AppId      string
	UserId     string
}

type GTplMallQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

type GTplMallQueryResult struct {
	Data       []*GTplMall
	PageResult *PaginationResult
}
