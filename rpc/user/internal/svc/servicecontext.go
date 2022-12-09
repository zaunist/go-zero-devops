package svc

import (
	_ "github.com/lib/pq"

	"go-zero-devops/rpc/model"
	"go-zero-devops/rpc/user/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.USERModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewUSERModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache),
	}
}
