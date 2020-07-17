package api

import (
	"fmt"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/bll"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/config"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/ginplus"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"time"
)

var AppSet = wire.NewSet(wire.Struct(new(App), "*"))

// 应用相关接口
type App struct {
	AppBll     bll.IApp
	MenuBll    bll.IMenu
	PageBll    bll.IPageManager
	SettingBll bll.ISetting
}

//初始化应用相关资源
func (a *App) Init(c *gin.Context) {

	var params schema.AppQueryParam
	if err := ginplus.ParseJSON(c, &params); err != nil {
		ginplus.ResError(c, err)
		return
	}
	appId := params.AppId
	//判断应用是否已经初始化
	exist, err := a.AppBll.Query(c, appId)
	if exist {
		ginplus.ResCustomError(c, fmt.Errorf("该应用已经初始化过，无须重新初始化！"))
		return
	}

	err = a.AppBll.Init(c, appId)
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	err = a.MenuBll.InitData(c, config.C.Menu.Data)
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	err = a.PageBll.InitData(c, config.C.Page.Data)
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}

	//更新平台信息
	err = a.SettingBll.Truncate(c)
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	var settings schema.Settings
	settings = append(settings, &schema.Setting{
		Key:       "platform_name",
		Value:     params.PlatformName,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})

	settings = append(settings, &schema.Setting{
		Key:       "platform_logo",
		Value:     params.PlatformLogo,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})
	for _, setting := range settings {
		err := a.SettingBll.Create(c, *setting)
		if err != nil {
			ginplus.ResCustomError(c, err)
			return
		}
	}

	ginplus.ResCustomSuccess(c, appId)
}

func (a *App) Query(c *gin.Context) {
	appId := c.Param("id")
	if len(appId) <= 0 {
		appId = ginplus.GetDefaultAppId();
	}
	exist, err := a.AppBll.Query(c, appId)
	if err != nil {
		ginplus.ResCustomError(c, err)
		return
	}
	ginplus.ResCustomSuccess(c, exist)
}
