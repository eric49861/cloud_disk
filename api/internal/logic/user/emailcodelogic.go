package user

import (
	"cloud-disk/api/internal/svc"
	"cloud-disk/api/internal/types"
	"cloud-disk/rpc/user/user"
	"context"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailCodeLogic {
	return &EmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// EmailCode 用户验证码接口
func (l *EmailCodeLogic) EmailCode(req *types.EmailCodeRequest) (resp *types.EmailCodeResponse, err error) {

	codeResponse, err := l.svcCtx.User.EmailCode(l.ctx, &user.EmailCodeRequest{
		Email: req.Email,
	})
	if err != nil {
		log.Println("rpc 调用验证码服务失败: ", err.Error())
		return nil, err
	}
	return &types.EmailCodeResponse{
		Msg: codeResponse.Msg,
	}, nil
}
