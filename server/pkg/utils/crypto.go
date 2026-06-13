package utils

import (
	"Helios/common"

	"golang.org/x/crypto/bcrypt"
)

// 密码加密
func PasswordEncrypt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		common.SystemLog.Error("密码加密异常：", err.Error())
		return "", err
	}
	return string(hashedPassword), nil
}

// 密码验证
func PasswordVerify(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
