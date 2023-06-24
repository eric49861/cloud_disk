package common

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	// CodeLength 验证码的长度
	CodeLength = 6
	// GrantCode 邮箱的授权码
	GrantCode = "xxxxxxxxxxx"
	// TokenExpiredTime Token的过期时间
	TokenExpiredTime = time.Minute * 10
	// RefreshTokenExpiredTime 刷新token的过期时间
	RefreshTokenExpiredTime = time.Minute * 11
	// CodeExpiredTime 验证码的过期时间
	CodeExpiredTime = time.Minute * 10
	// SecretKey token签发秘钥
	SecretKey = "cloud-disk"
)

type UserClaim struct {
	jwt.RegisteredClaims
	Name  string
	Email string
}
