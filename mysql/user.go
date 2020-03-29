package mysql

func CheckRepeat(username, email string) bool {
	row := checkRepeatUser.QueryRow(email, username)
	var id int64
	if err := row.Scan(&id); err != nil {
		return true
	}
	if id == 0 {
		return true
	}
	return false
}

func AddUser(username, email, password, token string) error {
	_, err := addUser.Exec(username, email, password, token)
	return err
}

func GetUser(username, email, password string) string {
	row := checkUser.QueryRow(username, email, password)
	var token string
	row.Scan(&token)
	return token
}

func GetUID(token string) int64 {
	row := getUserByToken.QueryRow(token)
	var id int64
	row.Scan(&id)
	return id
}
