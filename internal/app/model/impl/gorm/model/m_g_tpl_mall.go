package model

import (
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/model"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/model/impl/gorm/entity"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/tanjiancheng/gin-amis-admin/pkg/errors"
)

var _ model.IGTplMall = (*GTplMall)(nil)

// GTplMallSet 注入GTplMall
var GTplMallSet = wire.NewSet(wire.Struct(new(GTplMall), "*"), wire.Bind(new(model.IGTplMall), new(*GTplMall)))

// GTplMall
type GTplMall struct {
	DB *gorm.DB
}

func (a *GTplMall) getQueryOption(opts ...schema.GTplMallQueryOptions) schema.GTplMallQueryOptions {
	var opt schema.GTplMallQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *GTplMall) Query(ctx context.Context, params schema.GTplMallQueryParam, opts ...schema.GTplMallQueryOptions) (*schema.GTplMallQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetGTplMallDB(ctx, a.DB)
	if v := params.QueryValue; v != "" {
		v = "%" + v + "%"
		db = db.Where("name LIKE ?", v, v)
	}
	db = db.Where("scope LIKE ? or scope LIKE ?", params.AppId, "*")
	if len(params.UserId) > 0 {
		db = db.Where("status = ? or creator = ?", 1, params.UserId)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.GTplMalls
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.GTplMallQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaGTplMalls(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *GTplMall) Get(ctx context.Context, id string, opts ...schema.GTplMallQueryOptions) (*schema.GTplMall, error) {
	db := entity.GetGTplMallDB(ctx, a.DB).Where("id=?", id)
	var item entity.GTplMall
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaGTplMall(), nil
}

// Get 查询指定数据
func (a *GTplMall) GetByIdentify(ctx context.Context, identify string) (*schema.GTplMall, error) {
	db := entity.GetGTplMallDB(ctx, a.DB).Where("identify=?", identify)
	var item entity.GTplMall
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaGTplMall(), nil
}

// Create 创建数据
func (a *GTplMall) Create(ctx context.Context, item schema.GTplMall) error {
	eitem := entity.SchemaGTplMall(item).ToGTplMall()
	result := entity.GetGTplMallDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

// Update 更新数据
func (a *GTplMall) Update(ctx context.Context, id string, item schema.GTplMall) error {
	eitem := entity.SchemaGTplMall(item).ToGTplMall()
	result := entity.GetGTplMallDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

// Delete 删除数据
func (a *GTplMall) Delete(ctx context.Context, id string) error {
	result := entity.GetGTplMallDB(ctx, a.DB).Where("id=?", id).Delete(entity.GTplMall{})
	return errors.WithStack(result.Error)
}

// UpdateStatus 更新状态
func (a *GTplMall) UpdateStatus(ctx context.Context, id string, status int) error {
	result := entity.GetGTplMallDB(ctx, a.DB).Where("id=?", id).Update("status", status)
	return errors.WithStack(result.Error)
}
