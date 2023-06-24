package user

import (
	"cloud-disk/api/internal/svc"
	"cloud-disk/api/internal/types"
	"cloud-disk/rpc/user/user"
	"context"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login 登录接口
func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {

	loginResponse, err := l.svcCtx.User.Login(l.ctx, &user.LoginRequest{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		log.Println("rpc 调用登录服务失败: ", err.Error())
		return nil, err
	}

	return &types.LoginResponse{
		Token:        loginResponse.Token,
		RefreshToken: loginResponse.RefreshToken,
	}, nil
}
