package helper

import (
	"cloud-disk/rpc/user/common"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

// HashWithMD5 使用MD5算法获取字符串的hash值
func HashWithMD5(s string) string {
	resultByte := md5.Sum([]byte(s))
	result := fmt.Sprintf("%x", resultByte)
	return result
}

// EmailCode 向指定的邮箱发送验证码
func EmailCode(emailAddress, code string) error {
	e := email.NewEmail()
	e.From = "CLOUDDISK-ERIC <sang1585416826@163.com>"
	e.To = []string{emailAddress}
	e.Subject = "注册验证码"
	e.HTML = []byte("您的验证码 <h1>" + code + "</h1>, 有效时间为 " + strconv.Itoa(int(common.CodeExpiredTime.Minutes())) + " 分钟")
	err := e.SendWithTLS(
		"smtp.163.com:465",
		smtp.PlainAuth("", "sang1585416826@163.com", common.GrantCode, "smtp.163.com"),
		&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         "smtp.163.com",
		})
	if err != nil {
		return err
	}
	return nil
}

// RandCode 产生随机的验证码
func RandCode() string {
	nums := "0123456789"
	rand.Seed(time.Now().Unix())
	ans := make([]byte, 0)
	for i := 0; i < common.CodeLength; i++ {
		index := int(rand.Float64() * 10)
		ans = append(ans, nums[index])
	}
	return string(ans)
}

// UUID 生成一个随机的uuid
func UUID() string {
	return uuid.New().String()
}

// GenerateToken 根据用户信息生成token
func GenerateToken(name, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &common.UserClaim{
		Name:  name,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(common.TokenExpiredTime),
			},
			Issuer: "eric",
		},
	})
	tokenString, err := token.SignedString([]byte(common.SecretKey))
	return tokenString, err
}

// GenerateRefreshToken 刷新token
func GenerateRefreshToken(name, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &common.UserClaim{
		Name:  name,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(common.RefreshTokenExpiredTime),
			},
			Issuer: "eric",
		},
	})
	tokenString, err := token.SignedString([]byte(common.SecretKey))
	return tokenString, err
}
