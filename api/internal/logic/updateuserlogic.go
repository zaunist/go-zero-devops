package logic

import (
	"context"

	"go-zero-devops/api/internal/svc"
	"go-zero-devops/api/internal/types"
	"go-zero-devops/rpc/user/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.ReqUpdateUser) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.User.Update(l.ctx, &users.User{Id: req.Id, Password: req.Password, Username: req.Username})
	if err != nil {
		return nil, err
	}
	return &types.CommResp{Ok: r.Ok}, nil
}
