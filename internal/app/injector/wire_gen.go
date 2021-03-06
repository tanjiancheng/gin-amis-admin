// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injector

import (
	"github.com/tanjiancheng/gin-amis-admin/internal/app/api"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/bll/impl/bll"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/model/impl/gorm/model"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/module/adapter"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/router"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, func(), error) {
	auther, cleanup, err := InitAuth()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := InitGormDB()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	role := &model.Role{
		DB: db,
	}
	roleMenu := &model.RoleMenu{
		DB: db,
	}
	menuActionResource := &model.MenuActionResource{
		DB: db,
	}
	user := &model.User{
		DB: db,
	}
	userRole := &model.UserRole{
		DB: db,
	}
	casbinAdapter := &adapter.CasbinAdapter{
		RoleModel:         role,
		RoleMenuModel:     roleMenu,
		MenuResourceModel: menuActionResource,
		UserModel:         user,
		UserRoleModel:     userRole,
	}
	syncedEnforcer, cleanup3, err := InitCasbin(casbinAdapter)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	demo := &model.Demo{
		DB: db,
	}
	bllDemo := &bll.Demo{
		DemoModel: demo,
	}
	apiDemo := &api.Demo{
		DemoBll: bllDemo,
	}
	menu := &model.Menu{
		DB: db,
	}
	menuAction := &model.MenuAction{
		DB: db,
	}
	pageManagerModel := &model.PageManager{
		DB: db,
	}
	pageVersionHistoryModel := &model.PageVersionHistory{
		DB: db,
	}
	setting := &model.Setting{
		DB: db,
	}
	appModel := &model.App{
		DB: db,
	}
	gPlatformModel := &model.GPlatform{
		DB: db,
	}
	gTplModel := &model.GTplMall{
		DB: db,
	}
	login := &bll.Login{
		Enforcer:        syncedEnforcer,
		Auth:            auther,
		UserModel:       user,
		UserRoleModel:   userRole,
		RoleModel:       role,
		RoleMenuModel:   roleMenu,
		MenuModel:       menu,
		MenuActionModel: menuAction,
	}
	trans := &model.Trans{
		DB: db,
	}
	bllMenu := &bll.Menu{
		TransModel:              trans,
		MenuModel:               menu,
		MenuActionModel:         menuAction,
		MenuActionResourceModel: menuActionResource,
	}

	bllPage := &bll.PageManager{
		TransModel:              trans,
		PageManagerModel:        pageManagerModel,
		PageVersionHistoryModel: pageVersionHistoryModel,
	}
	bllPageVersionHistory := &bll.PageVersionHistory{
		TransModel:              trans,
		PageVersionHistoryModel: pageVersionHistoryModel,
	}
	bllSetting := &bll.Setting{
		TransModel:   trans,
		SettingModel: setting,
	}
	bllApp := &bll.App{
		AppModel: appModel,
	}
	bllGPlatformBll := &bll.GPlatform{
		GPlatformModel: gPlatformModel,
	}
	bllGTplMallBll := &bll.GTplMall{
		GTplMallModel: gTplModel,
		TransModel:    trans,
	}
	apiLogin := &api.Login{
		LoginBll:     login,
		GPlatformBll: bllGPlatformBll,
	}
	apiMenu := &api.Menu{
		MenuBll: bllMenu,
	}
	bllRole := &bll.Role{
		Enforcer:      syncedEnforcer,
		TransModel:    trans,
		RoleModel:     role,
		RoleMenuModel: roleMenu,
		UserModel:     user,
	}
	apiRole := &api.Role{
		RoleBll: bllRole,
	}
	bllUser := &bll.User{
		Enforcer:      syncedEnforcer,
		TransModel:    trans,
		UserModel:     user,
		UserRoleModel: userRole,
		RoleModel:     role,
	}
	apiUser := &api.User{
		UserBll: bllUser,
	}
	apiPageManager := &api.PageManager{
		PageManagerBll: bllPage,
		MenuBll:        bllMenu,
	}
	apiPageVersionHistory := &api.PageVersionHistory{
		PageVersionHistoryBll: bllPageVersionHistory,
	}
	apiSetting := &api.Setting{
		SettingBll:   bllSetting,
		GPlatformBll: bllGPlatformBll,
	}
	apiApp := &api.App{
		AppBll:       bllApp,
		MenuBll:      bllMenu,
		PageBll:      bllPage,
		SettingBll:   bllSetting,
		GPlatformBll: bllGPlatformBll,
	}
	apiGPlatform := &api.GPlatform{
		GPlatformBll: bllGPlatformBll,
	}
	apiGTplMall := &api.GTplMall{
		GTplMallBll: bllGTplMallBll,
	}
	routerRouter := &router.Router{
		Auth:                  auther,
		CasbinEnforcer:        syncedEnforcer,
		DemoAPI:               apiDemo,
		LoginAPI:              apiLogin,
		MenuAPI:               apiMenu,
		RoleAPI:               apiRole,
		UserAPI:               apiUser,
		PageManagerAPI:        apiPageManager,
		PageVersionHistoryAPI: apiPageVersionHistory,
		SettingAPI:            apiSetting,
		AppAPI:                apiApp,
		GPlatFormAPI:          apiGPlatform,
		GTplMallAPI:           apiGTplMall,
	}
	engine := InitGinEngine(routerRouter)
	injector := &Injector{
		Engine:         engine,
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		MenuBll:        bllMenu,
		PageBll:        bllPage,
		GPlatformBll:   bllGPlatformBll,
		GTplMallBll:    bllGTplMallBll,
	}
	return injector, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
