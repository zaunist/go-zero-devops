// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"go-zero-devops/rpc/user/internal/logic"
	"go-zero-devops/rpc/user/internal/svc"
	"go-zero-devops/rpc/user/types/user"
)

type UsersServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUsersServer
}

func NewUsersServer(svcCtx *svc.ServiceContext) *UsersServer {
	return &UsersServer{
		svcCtx: svcCtx,
	}
}

func (s *UsersServer) Login(ctx context.Context, in *user.ReqUser) (*user.RespLogin, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UsersServer) Create(ctx context.Context, in *user.ReqUser) (*user.CommResp, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *UsersServer) Update(ctx context.Context, in *user.User) (*user.CommResp, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *UsersServer) Delete(ctx context.Context, in *user.ReqUserId) (*user.CommResp, error) {
	l := logic.NewDeleteLogic(ctx, s.svcCtx)
	return l.Delete(in)
}

func (s *UsersServer) Get(ctx context.Context, in *user.ReqUserId) (*user.User, error) {
	l := logic.NewGetLogic(ctx, s.svcCtx)
	return l.Get(in)
}

func (s *UsersServer) GetAll(ctx context.Context, in *user.ReqGetAll) (*user.RespGetAll, error) {
	l := logic.NewGetAllLogic(ctx, s.svcCtx)
	return l.GetAll(in)
}
