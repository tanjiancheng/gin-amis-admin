package model

import (
	"context"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/ginplus"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/model"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/model/impl/gorm/entity"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var _ model.IApp = (*App)(nil)

var AppSet = wire.NewSet(wire.Struct(new(App), "*"), wire.Bind(new(model.IApp), new(*App)))

type App struct {
	DB *gorm.DB
}

// 创建对应的应用表
func (a *App) Init(ctx context.Context, appId string) error {
	ginplus.SetTablePrefix(appId)
	err := a.DB.AutoMigrate(
		new(entity.Demo),
		new(entity.MenuAction),
		new(entity.MenuActionResource),
		new(entity.Menu),
		new(entity.RoleMenu),
		new(entity.Role),
		new(entity.UserRole),
		new(entity.User),
		new(entity.PageManager),
		new(entity.PageVersionHistory),
		new(entity.Setting),
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Query(ctx context.Context, appId string) (bool, error) {
	ginplus.SetTablePrefix(appId)
	settingTableExists := a.DB.HasTable(new(entity.Setting))
	return settingTableExists, nil
}
