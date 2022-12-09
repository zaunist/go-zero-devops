package logic

import (
	"context"
	"fmt"

	"go-zero-devops/api/internal/svc"
	"go-zero-devops/api/internal/types"
	"go-zero-devops/rpc/user/users"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUsersLogic {
	return &GetAllUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllUsersLogic) GetAllUsers() (resp []types.User, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("l.svcCtx.User: ")
	r, err := l.svcCtx.User.GetAll(l.ctx, &users.ReqGetAll{})
	if err != nil {
		return nil, err
	}
	fmt.Println("r: ", r)
	res := make([]types.User, len(r.Users))
	for i := 0; i < len(r.Users); i++ {
		item := r.Users[i]
		res[i] = types.User{
			Id:       item.Id,
			Username: item.Username,
			Password: item.Password,
		}
	}
	return res, nil
}
