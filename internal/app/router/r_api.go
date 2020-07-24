package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/middleware"
)

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")

	g.Use(middleware.UserAuthMiddleware(a.Auth,
		middleware.AllowPathPrefixSkipper("/api/v1/pub/login"),
	))

	g.Use(middleware.CasbinMiddleware(a.CasbinEnforcer,
		middleware.AllowPathPrefixSkipper("/api/v1/pub"),
	))

	g.Use(middleware.RateLimiterMiddleware())

	g.Use(middleware.AppIdMiddleware())

	v1 := g.Group("/v1")
	{
		pub := v1.Group("/pub")
		{
			gLogin := pub.Group("login")
			{
				gLogin.GET("captchaid", a.LoginAPI.GetCaptcha)
				gLogin.GET("captcha", a.LoginAPI.ResCaptcha)
				gLogin.POST("", a.LoginAPI.Login)
				gLogin.POST("exit", a.LoginAPI.Logout)
			}

			gCurrent := pub.Group("current")
			{
				gCurrent.PUT("password", a.LoginAPI.UpdatePassword)
				gCurrent.GET("user", a.LoginAPI.GetUserInfo)
				gCurrent.GET("menutree", a.LoginAPI.QueryUserMenuTree)
			}

			gPubSetting := pub.Group("setting")
			{
				gPubSetting.GET("", a.SettingAPI.Query)
			}

			gPubApp := pub.Group("app")
			{
				gPubApp.GET(":id", a.AppAPI.Query)
			}

			gPageManager := pub.Group("page_manager")
			{
				gPageManager.GET(":id", a.PageManagerAPI.Get)
			}

			gPlatform := pub.Group("platforms")
			{
				gPlatform.GET(":id", a.GPlatFormAPI.GetOptions)
			}

			pub.POST("/refresh-token", a.LoginAPI.RefreshToken)
		}

		gDemo := v1.Group("demos")
		{
			gDemo.GET("", a.DemoAPI.Query)
			gDemo.GET(":id", a.DemoAPI.Get)
			gDemo.POST("", a.DemoAPI.Create)
			gDemo.PUT(":id", a.DemoAPI.Update)
			gDemo.DELETE(":id", a.DemoAPI.Delete)
			gDemo.PATCH(":id/enable", a.DemoAPI.Enable)
			gDemo.PATCH(":id/disable", a.DemoAPI.Disable)
		}

		gMenu := v1.Group("menus")
		{
			gMenu.GET("", a.MenuAPI.Query)
			gMenu.GET(":id", a.MenuAPI.Get)
			gMenu.POST("", a.MenuAPI.Create)
			gMenu.PUT(":id", a.MenuAPI.Update)
			gMenu.DELETE(":id", a.MenuAPI.Delete)
			gMenu.PATCH(":id/enable", a.MenuAPI.Enable)
			gMenu.PATCH(":id/disable", a.MenuAPI.Disable)
		}
		v1.GET("/menus.tree", a.MenuAPI.QueryTree)

		gRole := v1.Group("roles")
		{
			gRole.GET("", a.RoleAPI.Query)
			gRole.GET(":id", a.RoleAPI.Get)
			gRole.POST("", a.RoleAPI.Create)
			gRole.PUT(":id", a.RoleAPI.Update)
			gRole.DELETE(":id", a.RoleAPI.Delete)
			gRole.PATCH(":id/enable", a.RoleAPI.Enable)
			gRole.PATCH(":id/disable", a.RoleAPI.Disable)
		}
		v1.GET("/roles.select", a.RoleAPI.QuerySelect)

		gUser := v1.Group("users")
		{
			gUser.GET("", a.UserAPI.Query)
			gUser.GET(":id", a.UserAPI.Get)
			gUser.POST("", a.UserAPI.Create)
			gUser.PUT(":id", a.UserAPI.Update)
			gUser.DELETE(":id", a.UserAPI.Delete)
			gUser.PATCH(":id/enable", a.UserAPI.Enable)
			gUser.PATCH(":id/disable", a.UserAPI.Disable)
		}

		gPageManager := v1.Group("page_manager")
		{
			gPageManager.GET("", a.PageManagerAPI.Query)
			gPageManager.GET(":id", a.PageManagerAPI.Get)
			gPageManager.GET(":id/history", a.PageManagerAPI.History)
			gPageManager.POST("", a.PageManagerAPI.Create)
			gPageManager.POST(":id", a.PageManagerAPI.Clone)
			gPageManager.PUT(":id", a.PageManagerAPI.Update)
			gPageManager.DELETE(":id", a.PageManagerAPI.Delete)
			gPageManager.PATCH(":id/source", a.PageManagerAPI.Revert)
		}

		gPageVersionHistory := v1.Group("page_version_history")
		{
			gPageVersionHistory.GET(":id", a.PageVersionHistoryAPI.Get)
		}

		gSetting := v1.Group("setting")
		{
			gSetting.GET("", a.SettingAPI.Query)
			gSetting.POST("", a.SettingAPI.Create)
		}

		gApp := v1.Group("app")
		{
			gApp.POST("", a.AppAPI.Init)
			gApp.GET(":id", a.AppAPI.Query)
		}

		gPlatform := v1.Group("platforms")
		{
			gPlatform.GET("", a.GPlatFormAPI.Query)
			gPlatform.GET(":id", a.GPlatFormAPI.Get)
			gPlatform.POST("", a.GPlatFormAPI.Create)
			gPlatform.PUT(":id", a.GPlatFormAPI.Update)
			gPlatform.DELETE(":id", a.GPlatFormAPI.Delete)
			gPlatform.PATCH(":id/enable", a.GPlatFormAPI.Enable)
			gPlatform.PATCH(":id/disable", a.GPlatFormAPI.Disable)
		}

		gTplMall := v1.Group("tpl_mall")
		{
			gTplMall.GET("", a.GTplMallAPI.Query)
			gTplMall.GET(":id", a.GTplMallAPI.Get)
			gTplMall.POST("", a.GTplMallAPI.Create)
			gTplMall.PUT(":id", a.GTplMallAPI.Update)
			gTplMall.DELETE(":id", a.GTplMallAPI.Delete)
			gTplMall.PATCH(":id/publish", a.GTplMallAPI.Publish)
			gTplMall.PATCH(":id/enable", a.GTplMallAPI.Enable)
			gTplMall.PATCH(":id/disable", a.GTplMallAPI.Disable)
		}
		v1.Any("/tpl_mall.mock/:identity/:path", a.GTplMallAPI.Mock)      //用于预览里面的api接口
		v1.Any("/tpl_mall.mock/:identity/:path/:tid", a.GTplMallAPI.Mock) //用于预览里面的api接口
	}
}
