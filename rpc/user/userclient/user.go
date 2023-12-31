// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"cloud-disk/rpc/user/user"
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	EmailCodeRequest   = user.EmailCodeRequest
	EmailCodeResponse  = user.EmailCodeResponse
	LoginRequest       = user.LoginRequest
	LoginResponse      = user.LoginResponse
	RegisterRequest    = user.RegisterRequest
	RegisterResponse   = user.RegisterResponse
	UserDetailRequest  = user.UserDetailRequest
	UserDetailResponse = user.UserDetailResponse

	User interface {
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		UserDetail(ctx context.Context, in *UserDetailRequest, opts ...grpc.CallOption) (*UserDetailResponse, error)
		EmailCode(ctx context.Context, in *EmailCodeRequest, opts ...grpc.CallOption) (*EmailCodeResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) UserDetail(ctx context.Context, in *UserDetailRequest, opts ...grpc.CallOption) (*UserDetailResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserDetail(ctx, in, opts...)
}

func (m *defaultUser) EmailCode(ctx context.Context, in *EmailCodeRequest, opts ...grpc.CallOption) (*EmailCodeResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.EmailCode(ctx, in, opts...)
}
