package bll

import (
	"context"

	"github.com/tanjiancheng/gin-amis-admin/internal/app/bll"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/model"
	"github.com/google/wire"
)

var _ bll.IApp = (*App)(nil)

var AppSet = wire.NewSet(wire.Struct(new(App), "*"), wire.Bind(new(bll.IApp), new(*App)))

type App struct {
	AppModel model.IApp
}

func (a *App) Init(ctx context.Context, appId string) error {
	return a.AppModel.Init(ctx, appId)
}

func (a *App) Query(ctx context.Context, appId string) (bool, error) {
	return a.AppModel.Query(ctx, appId)
}
