package schema

type GPlatform struct {
	ID         int    `json:"id"`                        // 唯一标识
	AppID      string `json:"app_id" binding:"required"` // 平台ID
	IsCurrent  string `json:"is_current"`                // 是否当前平台
	Name       string `json:"name" binding:"required"`   // 平台名称
	Status     int    `json:"status"`                    // 状态(1:启用 -1:停用)
	CreateTime int64  `json:"create_time"`               // 创建时间
}

type GPlatformQueryParam struct {
	PaginationParam
	QueryValue string `form:"queryValue"` // 查询值
}

type GPlatformQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

type GPlatformQueryResult struct {
	Data       []*GPlatform
	PageResult *PaginationResult
}
