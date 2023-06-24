package user

import (
	"cloud-disk/api/internal/svc"
	"cloud-disk/api/internal/types"
	"cloud-disk/rpc/user/user"
	"context"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserDetail 用户详情接口
func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {

	detailResponse, err := l.svcCtx.User.UserDetail(l.ctx, &user.UserDetailRequest{
		UUID: req.UUID,
	})
	if err != nil {
		log.Println("rpc 调用用户详情服务失败: ", err.Error())
		return nil, err
	}

	return &types.UserDetailResponse{
		Name:  detailResponse.Name,
		Email: detailResponse.Email,
	}, nil
}
