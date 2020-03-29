package service

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/smoothlee/calendar/mysql"
)

func Register(username, email, pwd string) (code int, msg string, token string) {
	if mysql.CheckRepeat(username, email) {
		return -2, "user has already registered, please login", ""
	}
	pwdMd5 := md5V(pwd)
	token = md5V(username)
	if err := mysql.AddUser(username, email, pwdMd5, token); err != nil {
		return 0, err.Error(), ""
	}
	return 0, "OK", token
}

func Login(username, email, pwd string) (code int, msg string, uid string) {
	pwdMd5 := md5V(pwd)
	token := mysql.GetUser(username, email, pwdMd5)
	if token == "" {
		return -1, "please check your email, username and password", ""
	}
	return 0, "OK", token
}

func GetUIDByToken(token string) int64 {
	id := mysql.GetUID(token)
	return id
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
