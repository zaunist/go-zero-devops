package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"go-zero-devops/rpc/user/internal/svc"
	"go-zero-devops/rpc/user/types/user"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *user.ReqUserId) (*user.User, error) {
	// todo: add your logic here and delete this line

	one, err := l.svcCtx.Model.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &user.User{
		Id:       one.Id,
		Username: one.Username,
		Password: one.Password,
	}, nil
}
