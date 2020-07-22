package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/config"
	"github.com/tanjiancheng/gin-amis-admin/internal/app/ginplus"
	"strings"
)

// AppIdMiddleware 应用ID中间件，区分访问不同的应用请求
func AppIdMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		appID := ginplus.GetScopeAppId(c)
		ginplus.SetAppID(c, appID)
		c.Writer.Header().Set("X-App-Id", appID)
		ginplus.SetTablePrefix(appID) //切换表前缀

		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultName string) string {
			if len(defaultName) <= 0 {
				return defaultName
			}
			path := strings.Split(defaultName, "_")
			prefix := path[0]
			if prefix+"_" == config.C.Gorm.GlobalTablePrefix { //全局表不需要切换
				return defaultName
			}
			oldPath := prefix + "_" + path[1]
			newPath := prefix + "_" + appID
			return strings.Replace(defaultName, oldPath, newPath, -1)
		}
		c.Next()
	}
}
