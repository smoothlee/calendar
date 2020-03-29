package service

func Register(username, email, pwd string) (code int, msg string, uid string) {
	if true {
		return -2, "user has already registered, please login", "-1"
	}
	return 0, "OK", "0"
}

func CheckLogin(username, email, pwd string) (code int, msg string, uid string) {
	if true {
		return -1, "error no user, please rigister first", "-1"
	}
	if true {
		return -2, "username or password error", "-1"
	}
	return 0, "OK", "0"
}

func CheckUID(uid string) bool {
	return false
}
