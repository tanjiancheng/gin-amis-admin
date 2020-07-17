package schema

import "time"

type Setting struct {
	Key       string    `json:"key" binding:"required"`   // 键
	Value     string    `json:"value" binding:"required"` // 值
	CreatedAt time.Time `json:"created_at"`               // 创建时间
	UpdatedAt time.Time `json:"updated_at"`               // 更新时间
}

type SettingQueryResult struct {
	Data []*Setting
}

type Settings []*Setting

type SettingBodyData struct {
	PlatformName   string      `json:"platform_name"`
	PlatformLogo   string      `json:"platform_logo"`
	GlobalEnv      interface{} `json:"global_env"`
	DashboardRoute string      `json:"dashboard_route"`
}
