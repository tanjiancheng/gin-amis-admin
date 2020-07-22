package ginplus

import (
	"fmt"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/schema"
	"github.com/tanjiancheng/gin-amis-admin/pkg/errors"
	"github.com/tanjiancheng/gin-amis-admin/pkg/logger"
	"github.com/tanjiancheng/gin-amis-admin/pkg/util"
)

// 定义上下文中的键
const (
	prefix           = "gin-admin"
	UserIDKey        = prefix + "/user-id"
	ReqBodyKey       = prefix + "/req-body"
	ResBodyKey       = prefix + "/res-body"
	LoggerReqBodyKey = prefix + "/logger-req-body"
	AppIDKEY         = prefix + "/app-id"
)

//获取当前执行环境下的app_id
func GetScopeAppId(c *gin.Context) string {
	appID := c.GetHeader("X-App-Id")
	if appID == "" { //请求头获取不到从当前上下文获取
		appID = GetAppId(c)
	}

	if appID == "" {
		appID = GetDefaultAppId()
	}
	return appID
}

func GetDefaultAppId() string {
	return "default"
}

func GetDefaultAppName() string {
	return "后台系统"
}

var tablePrefix string

func SetTablePrefix(vals ...string) {
	newTablePrefix := config.C.Gorm.TablePrefix
	for _, val := range vals {
		newTablePrefix += val
		if val != config.C.Gorm.TablePrefix {
			newTablePrefix += "_"
		}
	}
	tablePrefix = newTablePrefix
}

func GetTableWithPrefix(defaultTableName string) string {
	if tablePrefix == "" {
		tablePrefix = config.C.Gorm.TablePrefix
	}
	return tablePrefix + defaultTableName
}

// GetToken 获取用户令牌
func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

// GetUserID 获取用户ID
func GetUserID(c *gin.Context) string {
	return c.GetString(UserIDKey)
}

// SetUserID 设定用户ID
func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDKey, userID)
}

func GetAppId(c *gin.Context) string {
	return c.GetString(AppIDKEY)
}

func SetAppID(c *gin.Context, appID string) {
	c.Set(AppIDKEY, appID)
}

// GetBody Get request body
func GetBody(c *gin.Context) []byte {
	if v, ok := c.Get(ReqBodyKey); ok {
		if b, ok := v.([]byte); ok {
			return b
		}
	}
	return nil
}

// ParseJSON 解析请求JSON
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	return nil
}

// ParseQuery 解析Query参数
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	return nil
}

// ParseForm 解析Form请求
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	return nil
}

// ResOK 响应OK
func ResOK(c *gin.Context) {
	ResSuccess(c, schema.StatusResult{Status: schema.OKStatus})
}

// ResList 响应列表数据
func ResList(c *gin.Context, v interface{}) {
	ResSuccess(c, schema.ListResult{List: v})
}

// ResPage 响应分页数据
func ResPage(c *gin.Context, v interface{}, pr *schema.PaginationResult) {
	list := schema.ListResult{
		List:       v,
		Pagination: pr,
	}
	ResSuccess(c, list)
}

//响应自定义成功
func ResCustomSuccess(c *gin.Context, v interface{}) {
	ResSuccess(c, &schema.ResponseSuccess{
		Status: 0,
		Msg:    schema.OKStatus.String(),
		Data:   v,
	})
}

//响应自定义失败
func ResCustomError(c *gin.Context, err error) {
	ResSuccess(c, &schema.ResponseSuccess{
		Status: -1,
		Msg:    err.Error(),
		Data:   nil,
	})
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ResJSON(c, http.StatusOK, v)
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := util.JSONMarshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}

// ResError 响应错误
func ResError(c *gin.Context, err error, status ...int) {
	ctx := c.Request.Context()
	var res *errors.ResponseError
	if err != nil {
		if e, ok := err.(*errors.ResponseError); ok {
			res = e
		} else {
			res = errors.UnWrapResponse(errors.Wrap500Response(err, "服务器错误"))
		}
	} else {
		res = errors.UnWrapResponse(errors.ErrInternalServer)
	}

	if len(status) > 0 {
		res.StatusCode = status[0]
	}

	if err := res.ERR; err != nil {
		if status := res.StatusCode; status >= 400 && status < 500 {
			logger.StartSpan(ctx).Warnf(err.Error())
		} else if status >= 500 {
			logger.ErrorStack(ctx, err)
		}
	}

	eitem := schema.ErrorItem{
		Code:    res.Code,
		Message: res.Message,
	}
	//ResJSON(c, res.StatusCode, schema.ErrorResult{Error: eitem})
	ResJSON(c, 200, schema.ErrorResult{Error: eitem})
}
