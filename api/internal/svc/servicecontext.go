package svc

import (
	"go-zero-devops/api/internal/config"
	"go-zero-devops/rpc/user/users"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	User   users.Users
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   users.NewUsers(zrpc.MustNewClient(c.User)),
	}
}
