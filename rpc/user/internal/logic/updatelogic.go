package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"go-zero-devops/rpc/model"
	"go-zero-devops/rpc/user/internal/svc"
	"go-zero-devops/rpc/user/types/user"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *user.User) (*user.CommResp, error) {
	// todo: add your logic here and delete this line

	err := l.svcCtx.Model.Update(l.ctx, &model.USER{
		Password: in.Password,
		Id:       in.Id,
		Username: in.Username,
	})
	if err != nil {
		return nil, err
	}
	return &user.CommResp{Ok: true}, nil
}
