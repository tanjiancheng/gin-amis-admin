package model

import (
	"context"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/model"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/model/impl/gorm/entity"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/tanjiancheng/gin-amis-admin/pkg/errors"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var _ model.IPageManager = (*PageManager)(nil)

// MenuSet 注入Menu
var PageManagerSet = wire.NewSet(wire.Struct(new(PageManager), "*"), wire.Bind(new(model.IPageManager), new(*PageManager)))

// Menu 菜单存储
type PageManager struct {
	DB *gorm.DB
}

func (a *PageManager) getQueryOption(opts ...schema.PageManagerQueryOptions) schema.PageManagerQueryOptions {
	var opt schema.PageManagerQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// 查询数据
func (a *PageManager) Query(ctx context.Context, params schema.PageManagerQueryParam, opts ...schema.PageManagerQueryOptions) (*schema.PageManagerQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetPageManagerDB(ctx, a.DB)

	if v := params.Name; v != "" {
		v = "%" + v + "%"
		db = db.Where(" name LIKE ?", v)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.PageManagers
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.PageManagerQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaPageManagers(),
	}
	return qr, nil
}

// 根据路由查询指定数据
func (a *PageManager) GetByRoute(ctx context.Context, route string, opts ...schema.PageManagerQueryOptions) (*schema.PageManager, error) {
	var item entity.PageManager
	ok, err := FindOne(ctx, entity.GetPageManagerDB(ctx, a.DB).Where("identify=?", route), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaPageManager(), nil
}

// 查询指定数据
func (a *PageManager) Get(ctx context.Context, id string, opts ...schema.PageManagerQueryOptions) (*schema.PageManager, error) {
	db := entity.GetPageManagerDB(ctx, a.DB).Where("id=?", id)
	var item entity.PageManager
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaPageManager(), nil
}

// 创建数据
func (a *PageManager) Create(ctx context.Context, item schema.PageManager) error {
	sitem := entity.SchemaPageManager(item)
	result := entity.GetPageManagerDB(ctx, a.DB).Create(sitem.ToPageManager())
	return errors.WithStack(result.Error)
}

// 更新数据
func (a *PageManager) Update(ctx context.Context, id string, item schema.PageManager) error {
	eitem := entity.SchemaPageManager(item).ToPageManager()
	result := entity.GetPageManagerDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

// 删除数据
func (a *PageManager) Delete(ctx context.Context, id string) error {
	result := entity.GetPageManagerDB(ctx, a.DB).Where("id=?", id).Delete(entity.PageManager{})
	return errors.WithStack(result.Error)
}

// 更新状态
func (a *PageManager) UpdateStatus(ctx context.Context, id string, status int) error {
	return nil
}

//获取最后一条数据
func (a *PageManager) GetLastId(ctx context.Context) (int, error) {
	var item entity.PageManager
	result := entity.GetPageManagerDB(ctx, a.DB).Order("id Desc").First(&item)
	if result.Error != nil {
		return 0, errors.WithStack(result.Error)
	}
	return item.ID, nil
}
