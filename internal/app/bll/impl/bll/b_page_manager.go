package bll

import (
	"context"
	"encoding/json"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/config"
	"github.com/tanjiancheng/gin-amis-admin/pkg/errors"
	"os"
	"strconv"
	"time"

	"github.com/google/wire"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/bll"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/model"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/tanjiancheng/gin-amis-admin/pkg/util"
)

var _ bll.IPageManager = (*PageManager)(nil)

// PageManagerSet 注入PageManager
var PageManagerSet = wire.NewSet(wire.Struct(new(PageManager), "*"), wire.Bind(new(bll.IPageManager), new(*PageManager)))

// 页面管理
type PageManager struct {
	TransModel              model.ITrans
	PageManagerModel        model.IPageManager
	PageVersionHistoryModel model.IPageVersionHistory
}

// InitData 初始化页面管理数据
func (a *PageManager) InitData(ctx context.Context, dataFile string) error {
	result, err := a.PageManagerModel.Query(ctx, schema.PageManagerQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		// 如果存在则不进行初始化
		return nil
	}

	data, err := a.readData(dataFile)
	if err != nil {
		return err
	}

	return a.createPages(ctx, data)
}

func (a *PageManager) readData(name string) (schema.PageManagers, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data schema.PageManagers
	d := util.YAMLNewDecoder(file)
	d.SetStrict(true)
	err = d.Decode(&data)
	return data, err
}

func (a *PageManager) createPages(ctx context.Context, list schema.PageManagers) error {
	return ExecTrans(ctx, a.TransModel, func(ctx context.Context) error {
		for _, item := range list {
			pitem := schema.PageManager{
				Identify:   item.Identify,
				Name:       item.Name,
				Source:     item.Source,
				Creator:    "root",
				CreateTime: time.Now().Unix(),
				UpdateTime: time.Now().Unix(),
			}

			err := a.Create(ctx, pitem)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// 查询数据
func (a *PageManager) Query(ctx context.Context, params schema.PageManagerQueryParam, opts ...schema.PageManagerQueryOptions) (*schema.PageManagerQueryResult, error) {
	return a.PageManagerModel.Query(ctx, params, opts...)
}

// 查询指定数据
func (a *PageManager) GetByRoute(ctx context.Context, route string, opts ...schema.PageManagerQueryOptions) (*schema.PageManager, error) {
	item, err := a.PageManagerModel.GetByRoute(ctx, route, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}
	return item, nil
}

// 查询指定数据
func (a *PageManager) Get(ctx context.Context, id string, opts ...schema.PageManagerQueryOptions) (*schema.PageManager, error) {
	item, err := a.PageManagerModel.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// 创建数据
func (a *PageManager) Create(ctx context.Context, item schema.PageManager) error {
	return a.PageManagerModel.Create(ctx, item)
}

// 更新数据
func (a *PageManager) Update(ctx context.Context, id string, item schema.PageManager) error {
	oldItem, err := a.PageManagerModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	item.ID = oldItem.ID
	item.UpdateTime = time.Now().Unix()

	//带有以下标识的页面不能更新
	cantNotUpdateIdentify := []string{
		"/tools/page_manager",
	}

	if config.C.RunMode != "debug" {
		for _, item := range cantNotUpdateIdentify {
			if item == oldItem.Identify {
				return errors.ErrNotAllowUpdate
			}
		}
	}

	err = ExecTrans(ctx, a.TransModel, func(ctx context.Context) error {
		pageMangerInfo, _ := json.Marshal(oldItem)
		pageVersionHistory := schema.PageVersionHistory{
			PageManagerId:   oldItem.ID,
			PageManagerInfo: string(pageMangerInfo),
			CreateTime:      time.Now().Unix(),
		}
		err := a.PageVersionHistoryModel.Create(ctx, pageVersionHistory)
		if err != nil {
			return err
		}
		return a.PageManagerModel.Update(ctx, id, item)
	})
	if err != nil {
		return err
	}
	return nil
}

// 删除数据
func (a *PageManager) Delete(ctx context.Context, id string) error {
	oldItem, err := a.PageManagerModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	//带有以下标识的页面不能删除(系统默认页面)
	cantNotDeleteIdentify := []string{
		"/tools/page_manager",
		"/system/user",
		"/system/role",
		"/system/menu",
		"/system/setting",
	}

	for _, item := range cantNotDeleteIdentify {
		if item == oldItem.Identify {
			return errors.ErrNotAllowDelete
		}
	}

	return a.PageManagerModel.Delete(ctx, id)
}

// 更新状态
func (a *PageManager) UpdateStatus(ctx context.Context, id string, status int) error {
	return nil
}

// 更新状态
func (a *PageManager) GetLastId(ctx context.Context) (int, error) {
	return a.PageManagerModel.GetLastId(ctx)
}

func (a *PageManager) History(ctx context.Context, params schema.PageVersionHistoryQueryParam, opts ...schema.PageVersionHistoryQueryOptions) (*schema.PageVersionHistoryQueryResult, error) {
	return a.PageVersionHistoryModel.Query(ctx, params, opts...)
}

// 更新数据
func (a *PageManager) Revert(ctx context.Context, pageVersionHistoryId string) error {
	pageVersionItem, err := a.PageVersionHistoryModel.Get(ctx, pageVersionHistoryId)
	if err != nil {
		return err
	} else if pageVersionItem == nil {
		return errors.ErrNotFound
	}
	pageManagerId := strconv.Itoa(pageVersionItem.PageManagerId)
	var versionPageManager schema.PageManager
	err = json.Unmarshal([]byte(pageVersionItem.PageManagerInfo), &versionPageManager)
	if err != nil {
		return err
	}
	var pageManager schema.PageManager
	pageManager.Name = versionPageManager.Name
	pageManager.Source = versionPageManager.Source
	pageManager.UpdateTime = time.Now().Unix()
	return a.PageManagerModel.Update(ctx, pageManagerId, pageManager)
}
