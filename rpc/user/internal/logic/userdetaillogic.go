package logic

import (
	"cloud-disk/rpc/user/internal/svc"
	"cloud-disk/rpc/user/model"
	"cloud-disk/rpc/user/user"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDetailLogic) UserDetail(in *user.UserDetailRequest) (*user.UserDetailResponse, error) {
	// 从数据库查询该用户
	u := &model.User{}
	err := l.svcCtx.SQL.Where("uuid = ?", in.UUID).First(u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("没有查到UUID = %s 的用户", in.UUID)
			return nil, errors.New("用户不存在")
		} else {
			log.Println("数据库查询失败, ", err.Error())
			return nil, errors.New("")
		}
	}
	return &user.UserDetailResponse{
		Name:  u.Name,
		Email: u.Email,
	}, nil
}
