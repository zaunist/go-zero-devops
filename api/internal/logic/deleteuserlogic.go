package logic

import (
	"context"

	"go-zero-devops/api/internal/svc"
	"go-zero-devops/api/internal/types"
	"go-zero-devops/rpc/user/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.ReqUserId) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.User.Delete(l.ctx, &users.ReqUserId{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.CommResp{Ok: r.Ok}, nil
}
