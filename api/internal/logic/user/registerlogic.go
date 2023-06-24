package user

import (
	"cloud-disk/api/internal/svc"
	"cloud-disk/api/internal/types"
	"cloud-disk/rpc/user/user"
	"context"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register 注册接口
func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	resp = new(types.RegisterResponse)
	response, err := l.svcCtx.User.Register(l.ctx, &user.RegisterRequest{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
		Code:     req.Code,
	})

	if err != nil {
		log.Println("rpc 调用注册服务失败: ", err.Error())
		return nil, err
	}
	return &types.RegisterResponse{
		Msg: response.Msg,
	}, nil
}
