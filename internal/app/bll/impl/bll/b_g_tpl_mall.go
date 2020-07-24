package bll

import (
	"context"
	"github.com/tanjiancheng/gin-amis-admin/pkg/util"
	"os"
	"regexp"
	"time"

	"github.com/google/wire"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/bll"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/model"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/tanjiancheng/gin-amis-admin/pkg/errors"
)

var _ bll.IGTplMall = (*GTplMall)(nil)

// GTplMallSet 注入GTplMall
var GTplMallSet = wire.NewSet(wire.Struct(new(GTplMall), "*"), wire.Bind(new(bll.IGTplMall), new(*GTplMall)))

// GTplMall 示例程序
type GTplMall struct {
	TransModel    model.ITrans
	GTplMallModel model.IGTplMall
}

func (a *GTplMall) InitData(ctx context.Context, dataFile string) error {
	result, err := a.GTplMallModel.Query(ctx, schema.GTplMallQueryParam{
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

	return a.createTplMalls(ctx, data)
}

func (a *GTplMall) readData(name string) (schema.GTplMalls, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data schema.GTplMalls
	d := util.YAMLNewDecoder(file)
	d.SetStrict(true)
	err = d.Decode(&data)
	return data, err
}

func (a *GTplMall) createTplMalls(ctx context.Context, list schema.GTplMalls) error {
	return ExecTrans(ctx, a.TransModel, func(ctx context.Context) error {
		for _, item := range list {
			pitem := schema.GTplMall{
				Identify:   item.Identify,
				Name:       item.Name,
				Desc:       item.Desc,
				Source:     item.Source,
				Status:     item.Status,
				Creator:    "root",
				Scope:      "*",
				Icon:       item.Icon,
				MockData:   item.MockData,
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

// Query 查询数据
func (a *GTplMall) Query(ctx context.Context, params schema.GTplMallQueryParam, opts ...schema.GTplMallQueryOptions) (*schema.GTplMallQueryResult, error) {
	return a.GTplMallModel.Query(ctx, params, opts...)
}

func (a *GTplMall) Check(ctx context.Context, appId string) error {
	if len(appId) <= 0 {
		return errors.NewResponse(-1, 200, "appId不合法")
	}
	r, _ := regexp.Compile("[A-Za-z0-9]")
	if !r.MatchString(appId) {
		return errors.NewResponse(-1, 200, "appId必须是[A-Za-z0-9]的格式")
	}
	//判断是否是数字或者字母
	return nil
}

// Get 查询指定数据
func (a *GTplMall) Get(ctx context.Context, id string, opts ...schema.GTplMallQueryOptions) (*schema.GTplMall, error) {
	item, err := a.GTplMallModel.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *GTplMall) GetByIdentify(ctx context.Context, identify string) (*schema.GTplMall, error) {
	item, err := a.GTplMallModel.GetByIdentify(ctx, identify)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Create 创建数据
func (a *GTplMall) Create(ctx context.Context, item schema.GTplMall) error {
	err := a.GTplMallModel.Create(ctx, item)
	if err != nil {
		return err
	}
	return nil
}

// Update 更新数据
func (a *GTplMall) Update(ctx context.Context, id string, item schema.GTplMall) error {
	oldItem, err := a.GTplMallModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	item.ID = oldItem.ID

	return a.GTplMallModel.Update(ctx, id, item)
}

// Delete 删除数据
func (a *GTplMall) Delete(ctx context.Context, id string) error {
	oldItem, err := a.GTplMallModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.GTplMallModel.Delete(ctx, id)
}

// UpdateStatus 更新状态
func (a *GTplMall) UpdateStatus(ctx context.Context, id string, status int) error {
	oldItem, err := a.GTplMallModel.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.GTplMallModel.UpdateStatus(ctx, id, status)
}
