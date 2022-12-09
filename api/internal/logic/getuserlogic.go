package logic

import (
	"context"

	"go-zero-devops/api/internal/svc"
	"go-zero-devops/api/internal/types"
	"go-zero-devops/rpc/user/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.ReqUserId) (resp *types.User, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.User.Get(l.ctx, &users.ReqUserId{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.User{
		Id:       r.Id,
		Username: r.Username,
		Password: r.Password,
	}, nil
}
