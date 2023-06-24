package helper

import (
	"testing"
	"time"
)

func TestHashWithMD5(t *testing.T) {
	s := "eric"
	hash := HashWithMD5(s)
	t.Logf("\"%s\" 的MD5值: %s", s, hash)
}

func TestUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		uuid := UUID()
		t.Logf("第 %d 次生成的随机uuid: %s", i+1, uuid)
		time.Sleep(time.Second * 1)
	}
}

func TestRandCode(t *testing.T) {
	for i := 0; i < 10; i++ {
		code := RandCode()
		t.Logf("第 %d 次生成的验证码: %s\n", i, code)
		time.Sleep(time.Second * 1)
	}
}

func TestEmailCode(t *testing.T) {
	err := EmailCode("1585416826@qq.com", RandCode())
	if err != nil {
		t.Fatal("验证码发送失败")
		return
	} else {
		t.Log("验证码发送成功")
	}
}

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("eric", "1585416826@qq.com")
	if err != nil {
		t.Fatal("token生成失败: ", err.Error())
	}
	t.Log(token)
}
