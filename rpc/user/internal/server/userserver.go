// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"cloud-disk/rpc/user/internal/logic"
	"cloud-disk/rpc/user/internal/svc"
	"cloud-disk/rpc/user/user"
	"context"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) UserDetail(ctx context.Context, in *user.UserDetailRequest) (*user.UserDetailResponse, error) {
	l := logic.NewUserDetailLogic(ctx, s.svcCtx)
	return l.UserDetail(in)
}

func (s *UserServer) EmailCode(ctx context.Context, in *user.EmailCodeRequest) (*user.EmailCodeResponse, error) {
	l := logic.NewEmailCodeLogic(ctx, s.svcCtx)
	return l.EmailCode(in)
}
