package logic

import (
	"cloud-disk/rpc/user/helper"
	"cloud-disk/rpc/user/internal/svc"
	"cloud-disk/rpc/user/model"
	"cloud-disk/rpc/user/user"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login 用户的登录逻辑
func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {

	u := new(model.User)
	err := l.svcCtx.SQL.Where("(name = ? OR email = ?) AND password = ?", in.Name, in.Name, helper.HashWithMD5(in.Password)).First(u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在")
	}
	if err != nil {
		log.Println("数据查询用户失败, ", err.Error())
		return nil, errors.New("登录失败")
	}
	tokenString, err := helper.GenerateToken(u.Name, u.Email)
	if err != nil {
		log.Println("token生成失败, ", err.Error())
		return nil, errors.New("登录失败")
	}
	refreshTokenString, err := helper.GenerateRefreshToken(u.Name, u.Email)
	if err != nil {
		log.Println("refreshToken生成失败, ", err.Error())
		return nil, errors.New("登录失败")
	}
	return &user.LoginResponse{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
