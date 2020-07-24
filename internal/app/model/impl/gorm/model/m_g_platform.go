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

var _ model.IGPlatform = (*GPlatform)(nil)

// GPlatformSet 注入GPlatform
var GPlatformSet = wire.NewSet(wire.Struct(new(GPlatform), "*"), wire.Bind(new(model.IGPlatform), new(*GPlatform)))

// GPlatform 示例存储
type GPlatform struct {
	DB *gorm.DB
}

func (a *GPlatform) getQueryOption(opts ...schema.GPlatformQueryOptions) schema.GPlatformQueryOptions {
	var opt schema.GPlatformQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *GPlatform) GetOptions(ctx context.Context) (*schema.GPlatformSelectOptions, error) {
	db := entity.GetGPlatformDB(ctx, a.DB)
	db = db.Where("status = ?", 1)
	var items entity.GPlatforms
	err := db.Find(&items).Error
	if err != nil {
		return nil, err
	}
	var selectOptions schema.GPlatformSelectOptions
	selectOptions.Options = append(selectOptions.Options, &schema.Option{
		Label: "全部",
		Value: "*",
	})
	for _, item := range items {
		selectOptions.Options = append(selectOptions.Options, &schema.Option{
			Label: item.Name,
			Value: item.AppID,
		})
	}
	return &selectOptions, nil
}

// Query 查询数据
func (a *GPlatform) Query(ctx context.Context, params schema.GPlatformQueryParam, opts ...schema.GPlatformQueryOptions) (*schema.GPlatformQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := entity.GetGPlatformDB(ctx, a.DB)
	if v := params.QueryValue; v != "" {
		v = "%" + v + "%"
		db = db.Where("name LIKE ? OR app_id LIKE ?", v, v)
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	db = db.Order(ParseOrder(opt.OrderFields))

	var list entity.GPlatforms
	pr, err := WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.GPlatformQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaGPlatforms(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *GPlatform) Get(ctx context.Context, id string, opts ...schema.GPlatformQueryOptions) (*schema.GPlatform, error) {
	db := entity.GetGPlatformDB(ctx, a.DB).Where("id=?", id)
	var item entity.GPlatform
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaGPlatform(), nil
}

// Get 查询指定数据
func (a *GPlatform) GetByAppId(ctx context.Context, appId string) (*schema.GPlatform, error) {
	db := entity.GetGPlatformDB(ctx, a.DB).Where("app_id=?", appId)
	var item entity.GPlatform
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaGPlatform(), nil
}

// Create 创建数据
func (a *GPlatform) Create(ctx context.Context, item schema.GPlatform) error {
	eitem := entity.SchemaGPlatform(item).ToGPlatform()
	result := entity.GetGPlatformDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

// Update 更新数据
func (a *GPlatform) Update(ctx context.Context, id string, item schema.GPlatform) error {
	eitem := entity.SchemaGPlatform(item).ToGPlatform()
	result := entity.GetGPlatformDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

// Delete 删除数据
func (a *GPlatform) Delete(ctx context.Context, id string) error {
	result := entity.GetGPlatformDB(ctx, a.DB).Where("id=?", id).Delete(entity.GPlatform{})
	return errors.WithStack(result.Error)
}

// UpdateStatus 更新状态
func (a *GPlatform) UpdateStatus(ctx context.Context, id string, status int) error {
	result := entity.GetGPlatformDB(ctx, a.DB).Where("id=?", id).Update("status", status)
	return errors.WithStack(result.Error)
}
