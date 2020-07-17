package iutil

import (
	"github.com/tanjiancheng/gin-amis-admin/internal/app/config"
	"github.com/tanjiancheng/gin-amis-admin/pkg/logger"
	"github.com/tanjiancheng/gin-amis-admin/pkg/trace"
	"github.com/tanjiancheng/gin-amis-admin/pkg/unique"
)

var idFunc = func() string {
	return unique.NewSnowflakeID().String()
}

// InitID ...
func InitID() {
	switch config.C.UniqueID.Type {
	case "uuid":
		idFunc = func() string {
			return unique.MustUUID().String()
		}
	case "object":
		idFunc = func() string {
			return unique.NewObjectID().Hex()
		}
	default:
		// Initialize snowflake node
		err := unique.SetSnowflakeNode(config.C.UniqueID.Snowflake.Node, config.C.UniqueID.Snowflake.Epoch)
		if err != nil {
			panic(err)
		}

		logger.SetTraceIDFunc(func() string {
			return unique.NewSnowflakeID().String()
		})

		trace.SetIDFunc(func() string {
			return unique.NewSnowflakeID().String()
		})

		idFunc = func() string {
			return unique.NewSnowflakeID().String()
		}
	}
}

// NewID Create unique id
func NewID() string {
	return idFunc()
}
