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

var _ model.IPageVersionHistory = (*PageVersionHistory)(nil)

// MenuSet 注入Menu
var PageVersionHistorySet = wire.NewSet(wire.Struct(new(PageVersionHistory), "*"), wire.Bind(new(model.IPageVersionHistory), new(*PageVersionHistory)))

// Menu 菜单存储
type PageVersionHistory struct {
	DB *gorm.DB
}

func (a *PageVersionHistory) getQueryOption(opts ...schema.PageVersionHistoryQueryOptions) schema.PageVersionHistoryQueryOptions {
	var opt schema.PageVersionHistoryQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// 查询数据
func (a *PageVersionHistory) Query(ctx context.Context, params schema.PageVersionHistoryQueryParam, opts ...schema.PageVersionHistoryQueryOptions) (*schema.PageVersionHistoryQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetPageVersionHistoryDB(ctx, a.DB)

	if v := params.PageManagerId; v != "" {
		db = db.Where(" page_manager_id = ?", v)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.PageVersionHistorys
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.PageVersionHistoryQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaPageVersionHistorys(),
	}
	return qr, nil
}

// 查询指定数据
func (a *PageVersionHistory) Get(ctx context.Context, id string, opts ...schema.PageVersionHistoryQueryOptions) (*schema.PageVersionHistory, error) {
	db := entity.GetPageVersionHistoryDB(ctx, a.DB).Where("id=?", id)
	var item entity.PageVersionHistory
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaPageVersionHistory(), nil
}

// 创建数据
func (a *PageVersionHistory) Create(ctx context.Context, item schema.PageVersionHistory) error {
	sitem := entity.SchemaPageVersionHistory(item)
	result := entity.GetPageVersionHistoryDB(ctx, a.DB).Create(sitem.ToPageVersionHistory())
	return errors.WithStack(result.Error)
}

// 更新数据
func (a *PageVersionHistory) Update(ctx context.Context, id string, item schema.PageVersionHistory) error {
	eitem := entity.SchemaPageVersionHistory(item).ToPageVersionHistory()
	result := entity.GetPageVersionHistoryDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

// 删除数据
func (a *PageVersionHistory) Delete(ctx context.Context, id string) error {
	result := entity.GetPageVersionHistoryDB(ctx, a.DB).Where("id=?", id).Delete(entity.PageVersionHistory{})
	return errors.WithStack(result.Error)
}
