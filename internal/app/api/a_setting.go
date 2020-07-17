package api

import (
	"encoding/json"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/bll"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/ginplus"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"time"
)

// SettingSet 注入Setting
var SettingSet = wire.NewSet(wire.Struct(new(Setting), "*"))

// Setting 示例程序
type Setting struct {
	SettingBll bll.ISetting
}

// Query 查询数据
func (a *Setting) Query(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := a.SettingBll.Query(ctx)
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}

	returnData := make(map[string]interface{})
	resultData := result.Data
	for _, item := range resultData {
		//判断是否json数据
		var jsonValue []map[string]interface{}
		if json.Unmarshal([]byte(item.Value), &jsonValue) == nil { //判断是否json
			returnData[item.Key] = jsonValue
		} else {
			returnData[item.Key] = item.Value
		}
	}

	ginplus.ResCustomSuccess(c, returnData)
}

// Get 查询指定数据
func (a *Setting) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.SettingBll.Get(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	ginplus.ResSuccess(c, item)
}

// Create 创建数据
func (a *Setting) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var settings schema.Settings
	var bodyData schema.SettingBodyData
	if err := ginplus.ParseJSON(c, &bodyData); err != nil {
		ginplus.ResCustomError(c, err)
		return
	}

	settings = append(settings, &schema.Setting{
		Key:       "platform_name",
		Value:     bodyData.PlatformName,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})

	settings = append(settings, &schema.Setting{
		Key:       "platform_logo",
		Value:     bodyData.PlatformLogo,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})

	settings = append(settings, &schema.Setting{
		Key:       "dashboard_route",
		Value:     bodyData.DashboardRoute,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})

	globalEnvStr, _ := json.Marshal(bodyData.GlobalEnv)
	settings = append(settings, &schema.Setting{
		Key:       "global_env",
		Value:     string(globalEnvStr),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})

	err := a.SettingBll.Truncate(ctx)
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	for _, setting := range settings {
		err := a.SettingBll.Create(ctx, *setting)
		if err != nil {
			ginplus.ResCustomError(c, err)
			return
		}
	}

	ginplus.ResCustomSuccess(c, nil)
}

// Update 更新数据
func (a *Setting) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Setting
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResCustomError(c, err)
		return
	}

	err := a.SettingBll.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Delete 删除数据
func (a *Setting) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.SettingBll.Delete(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	ginplus.ResOK(c)
}
