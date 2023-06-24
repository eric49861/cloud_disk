package logic

import (
	"cloud-disk/rpc/user/common"
	"cloud-disk/rpc/user/helper"
	"cloud-disk/rpc/user/internal/svc"
	"cloud-disk/rpc/user/user"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailCodeLogic {
	return &EmailCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// EmailCode 获取验证码逻辑
func (l *EmailCodeLogic) EmailCode(in *user.EmailCodeRequest) (*user.EmailCodeResponse, error) {
	// 检查当前redis中是否存在该邮箱的验证码
	_, err := l.svcCtx.RDB.Get(l.ctx, in.Email).Result()
	if err == redis.Nil {
		// 说明没有获取过验证码 or 验证码已过期
		code := helper.RandCode()
		err = helper.EmailCode(in.Email, code)
		if err != nil {
			log.Println("获取验证码出错, ", err)
			return nil, errors.New("获取验证码出错")
		}
		// 将验证码存入redis
		l.svcCtx.RDB.Set(l.ctx, in.Email, code, common.CodeExpiredTime)
	} else {
		return nil, errors.New("频繁获取验证码")
	}
	return &user.EmailCodeResponse{
		Msg: "验证码已发送,注意查收",
	}, nil
}
