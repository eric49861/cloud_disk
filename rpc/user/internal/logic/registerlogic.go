package logic

import (
	"cloud-disk/rpc/user/helper"
	"cloud-disk/rpc/user/internal/svc"
	"cloud-disk/rpc/user/model"
	"cloud-disk/rpc/user/user"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register 用户的注册逻辑
func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 查询是否存在已经存在的用户
	var count int64 = 0
	err := l.svcCtx.SQL.Table("user").Where("name = ? OR email = ?", in.Name, in.Email).Count(&count).Error
	if count > 0 {
		// 查询到用户
		return &user.RegisterResponse{
			Msg: "用户名或者邮箱已注册",
		}, nil
	} else {
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Register 错误, ", err)
			return nil, errors.New("因为未知原因, 注册失败, 请稍后重试")
		} else {
			// 检查验证码是否正确
			code, err := l.svcCtx.RDB.Get(l.ctx, in.Email).Result()
			if err == redis.Nil {
				return nil, errors.New("验证码已过期")
			}
			if code != in.Code {
				return nil, errors.New("验证码不正确")
			}
			//保存用户信息
			u := &model.User{
				UUID:     helper.UUID(),
				Name:     in.Name,
				Password: helper.HashWithMD5(in.Password),
				Email:    in.Email,
			}
			err = l.svcCtx.SQL.Create(u).Error
			if err != nil {
				log.Println("数据库创建用户记录出错, ", err)
				return nil, errors.New("因为未知原因,注册失败,请稍后重试")
			}
			return &user.RegisterResponse{
				Msg: "注册成功",
			}, nil
		}
	}
}
