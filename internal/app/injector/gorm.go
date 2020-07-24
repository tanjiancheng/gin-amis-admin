package injector

import (
	"context"
	"errors"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/ginplus"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"os"
	"path/filepath"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/config"
	igorm "github.com/tanjiancheng/gin-amis-admin/internal/app/model/impl/gorm"
)

// InitGormDB 初始化gorm存储
func InitGormDB() (*gorm.DB, func(), error) {
	cfg := config.C.Gorm
	db, cleanFunc, err := NewGormDB()
	if err != nil {
		return nil, cleanFunc, err
	}

	if cfg.EnableAutoMigrate {
		err = igorm.AutoMigrate(db)
		if err != nil {
			return nil, cleanFunc, err
		}
	}

	return db, cleanFunc, nil
}

func InitGormData(ctx context.Context, injector *Injector) error {
	// 初始化菜单数据
	if config.C.Menu.Enable && config.C.Menu.Data != "" {
		err := injector.MenuBll.InitData(ctx, config.C.Menu.Data)
		if err != nil {
			return err
		}
	}

	// 初始化页面管理数据
	if config.C.Page.Enable && config.C.Page.Data != "" {
		err := injector.PageBll.InitData(ctx, config.C.Page.Data)
		if err != nil {
			return err
		}
	}

	// 初始化模板商城数据
	if config.C.TplMall.Enable && config.C.TplMall.Data != "" {
		err := injector.GTplMallBll.InitData(ctx, config.C.TplMall.Data)
		if err != nil {
			return err
		}
	}

	//初始化平台数据
	err := injector.GPlatformBll.Create(ctx, schema.GPlatform{
		AppID:      ginplus.GetDefaultAppId(),
		Name:       ginplus.GetDefaultAppName(),
		Status:     1,
		CreateTime: time.Now().Unix(),
	})
	if err != nil {
		return err
	}
	return nil
}

// NewGormDB 创建DB实例
func NewGormDB() (*gorm.DB, func(), error) {
	cfg := config.C
	var dsn string
	switch cfg.Gorm.DBType {
	case "mysql":
		dsn = cfg.MySQL.DSN()
	case "sqlite3":
		dsn = cfg.Sqlite3.DSN()
		_ = os.MkdirAll(filepath.Dir(dsn), 0777)
	case "postgres":
		dsn = cfg.Postgres.DSN()
	default:
		return nil, nil, errors.New("unknown db")
	}

	return igorm.NewDB(&igorm.Config{
		Debug:        cfg.Gorm.Debug,
		DBType:       cfg.Gorm.DBType,
		DSN:          dsn,
		MaxIdleConns: cfg.Gorm.MaxIdleConns,
		MaxLifetime:  cfg.Gorm.MaxLifetime,
		MaxOpenConns: cfg.Gorm.MaxOpenConns,
	})
}
