package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/bll"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/ginplus"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
)

var GPlatFormSet = wire.NewSet(wire.Struct(new(GPlatform), "*"))

type GPlatform struct {
	GPlatformBll bll.IGPlatform
}

// Query 查询数据
func (a *GPlatform) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.GPlatformQueryParam
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.GPlatformBll.Query(ctx, params)

	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	currentAppId := ginplus.GetAppId(c)
	data := result.Data
	for _, item := range data {
		if item.AppID == currentAppId {
			item.IsCurrent = fmt.Sprintf("%s", "<i class='fa fa-check text-info'></i>")
		}
	}
	ginplus.ResPage(c, data, result.PageResult)
}

// Get 查询指定数据
func (a *GPlatform) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if id == "check" {
		appId := c.DefaultQuery("app_id", ginplus.GetDefaultAppId())
		err := a.GPlatformBll.Check(ctx, appId)
		if err != nil {
			ginplus.ResError(c, err)
			return
		}
		ginplus.ResCustomSuccess(c, nil)
		return
	}
	item, err := a.GPlatformBll.Get(ctx, id)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, item)
}

// Create 创建数据
func (a *GPlatform) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.GPlatform
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	err := a.GPlatformBll.Create(ctx, item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nil)
}

// Update 更新数据
func (a *GPlatform) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.GPlatform
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}
	err := a.GPlatformBll.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Delete 删除数据
func (a *GPlatform) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.GPlatformBll.Delete(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Enable 启用数据
func (a *GPlatform) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.GPlatformBll.UpdateStatus(ctx, c.Param("id"), 1)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Disable 禁用数据
func (a *GPlatform) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.GPlatformBll.UpdateStatus(ctx, c.Param("id"), -1)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
