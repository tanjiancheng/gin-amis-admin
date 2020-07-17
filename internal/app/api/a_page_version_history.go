package api

import (
	"encoding/json"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/bll"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/ginplus"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// PageVersionHistorySet 注入PageVersionHistory
var PageVersionHistorySet = wire.NewSet(wire.Struct(new(PageVersionHistory), "*"))

// PageVersionHistory 用户管理
type PageVersionHistory struct {
	PageVersionHistoryBll bll.IPageVersionHistory
}

// Query 查询数据
func (a *PageVersionHistory) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PageVersionHistoryQueryParam
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}
	params.Pagination = true
	result, err := a.PageVersionHistoryBll.Query(ctx, params)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	listData := map[string]interface{}{
		"rows":  result.Data,
		"count": result.PageResult.Total,
	}
	ginplus.ResCustomSuccess(c, listData)
}


// Get 查询指定数据
func (a *PageVersionHistory) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.PageVersionHistoryBll.Get(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	var pageManager schema.PageManager
	err = json.Unmarshal([]byte(item.PageManagerInfo), &pageManager)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	item.Name = pageManager.Name
	item.PageCreateTime = pageManager.CreateTime
	ginplus.ResCustomSuccess(c, item)
}
