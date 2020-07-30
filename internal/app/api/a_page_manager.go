package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/bll"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/ginplus"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"strconv"
	"time"
)

// PageManagerSet 注入PageManager
var PageManagerSet = wire.NewSet(wire.Struct(new(PageManager), "*"))

// PageManager 用户管理
type PageManager struct {
	PageManagerBll bll.IPageManager
	MenuBll        bll.IMenu
}

// Query 查询数据
func (a *PageManager) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PageManagerQueryParam
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}
	params.Pagination = true
	result, err := a.PageManagerBll.Query(ctx, params)
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

// GetByRoute 根据路由查询指定数据
func (a *PageManager) GetByRoute(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PageManagerQueryParam
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	item, err := a.PageManagerBll.GetByRoute(ctx, params.Route)
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	source := item.RenderSource
	sourceMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(source), &sourceMap)
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	ginplus.ResCustomSuccess(c, sourceMap)
}

// GetByRouteWithDetail 根据路由查询详细数据
func (a *PageManager) GetByRouteWithDetail(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PageManagerQueryParam
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	item, err := a.PageManagerBll.GetByRoute(ctx, params.Route)
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	source := item.Source
	sourceMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(source), &sourceMap)
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	ginplus.ResCustomSuccess(c, item)
}

// Get 查询指定数据
func (a *PageManager) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if id == "route" {
		a.GetByRoute(c)
		return
	}
	if id == "route_with_detail" {
		a.GetByRouteWithDetail(c)
		return
	}
	item, err := a.PageManagerBll.Get(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResCustomSuccess(c, item)
}

// Clone 克隆当前数据
func (a *PageManager) Clone(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	item, err := a.PageManagerBll.Get(ctx, id)

	//lastId 查询最后一条记录的id
	lastId, err := a.PageManagerBll.GetLastId(ctx)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	lastId++
	newIdentify := item.Identify + "_" + strconv.Itoa(lastId)
	err = a.PageManagerBll.Create(ctx, schema.PageManager{
		Identify:   newIdentify,
		Name:       item.Name,
		Source:     item.Source,
		Creator:    ginplus.GetUserID(c),
		CreateTime: time.Now().Unix(),
	})

	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResCustomSuccess(c, nil)
}

// Create 创建数据
func (a *PageManager) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.PageManager
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}
	item.Creator = ginplus.GetUserID(c)
	item.CreateTime = time.Now().Unix()
	err := a.PageManagerBll.Create(ctx, item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResCustomSuccess(c, nil)
}

// Update 更新数据
func (a *PageManager) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.PageManager
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	err := a.PageManagerBll.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResCustomSuccess(c, nil)
}

// Delete 删除数据
func (a *PageManager) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	item, err := a.PageManagerBll.Get(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	router := item.Identify
	menuInfo, err := a.MenuBll.GetByRouter(ctx, router)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	//如果有菜单数据先删除菜单
	if menuInfo != nil {
		err = a.MenuBll.Delete(ctx, menuInfo.ID)
		if err != nil {
			ginplus.ResError(c, err)
			return
		}
	}

	err = a.PageManagerBll.Delete(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResCustomSuccess(c, nil)
}

func (a *PageManager) History(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.PageVersionHistoryQueryParam
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}
	pageManagerId := c.Param("id")
	if len(pageManagerId) <= 0 {
		ginplus.ResError(c, fmt.Errorf("id为空"))
		return
	}
	params.PageManagerId = pageManagerId
	params.Pagination = true
	result, err := a.PageManagerBll.History(ctx, params)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	rows := result.Data
	for _, item := range rows {
		var pageManager schema.PageManager
		err := json.Unmarshal([]byte(item.PageManagerInfo), &pageManager)
		if err != nil {
			ginplus.ResError(c, err)
			return
		}
		item.Name = pageManager.Name
		item.PageCreateTime = pageManager.CreateTime
	}

	listData := map[string]interface{}{
		"rows":  result.Data,
		"count": result.PageResult.Total,
	}
	ginplus.ResCustomSuccess(c, listData)
}

func (a *PageManager) Revert(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.PageManagerBll.Revert(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResCustomSuccess(c, []string{})
}
