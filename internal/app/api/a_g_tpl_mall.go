package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/bll"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/ginplus"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"time"
)

var GTplMallSet = wire.NewSet(wire.Struct(new(GTplMall), "*"))

type GTplMall struct {
	GTplMallBll bll.IGTplMall
}

// Query 查询数据
func (a *GTplMall) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.GTplMallQueryParam
	if err := ginplus.ParseQuery(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}

	params.Pagination = true
	params.AppId = ginplus.GetAppId(c)
	params.UserId = ginplus.GetUserID(c)
	result, err := a.GTplMallBll.Query(ctx, params)

	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	data := result.Data
	for _, item := range data {
		item.IconFull = fmt.Sprintf("<i class='%s text-3x pull-left thumb b-3x m-r'></i>", item.Icon)
	}
	ginplus.ResPage(c, data, result.PageResult)
}

// Get 查询指定数据
func (a *GTplMall) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	item, err := a.GTplMallBll.Get(ctx, id)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, item)
}

func (a *GTplMall) Mock(c *gin.Context) {
	ctx := c.Request.Context()
	path := c.Param("path")
	identify := c.Param("identity")
	item, err := a.GTplMallBll.GetByIdentify(ctx, identify)
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	var mockRespnse map[string]interface{}
	mockData := item.MockData
	for _, item := range mockData {
		if item.Path == path {
			_ = json.Unmarshal([]byte(item.Data), &mockRespnse)
		}
	}
	ginplus.ResSuccess(c, mockRespnse)
	return
}

// Create 创建数据
func (a *GTplMall) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.GTplMall
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	item.Creator = ginplus.GetUserID(c)
	item.CreateTime = time.Now().Unix()
	err := a.GTplMallBll.Create(ctx, item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Update 更新数据
func (a *GTplMall) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.GTplMall
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}
	err := a.GTplMallBll.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

// Delete 删除数据
func (a *GTplMall) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.GTplMallBll.Delete(ctx, c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

func (a *GTplMall) Publish(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.GTplMallBll.UpdateStatus(ctx, c.Param("id"), 1)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

func (a *GTplMall) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.GTplMallBll.UpdateStatus(ctx, c.Param("id"), 1)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}

func (a *GTplMall) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.GTplMallBll.UpdateStatus(ctx, c.Param("id"), -1)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
