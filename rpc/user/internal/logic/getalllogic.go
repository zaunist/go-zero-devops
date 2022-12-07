package logic

import (
	"context"

	"go-zero-devops/rpc/user/internal/svc"
	"go-zero-devops/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllLogic {
	return &GetAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllLogic) GetAll(in *user.ReqGetAll) (*user.RespGetAll, error) {
	// todo: add your logic here and delete this line

	return &user.RespGetAll{}, nil
}
