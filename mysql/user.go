package mysql

import "fmt"

func CheckRepeat(username, email string) bool {
	row := checkRepeatUser.QueryRow(email, username)
	var id int64
	row.Scan(&id)
	if id != 0 {
		fmt.Println(id)
		return true
	}
	return false
}

func AddUser(username, email, password, token string) error {
	_, err := addUser.Exec(username, email, password, token)
	return err
}

func GetUser(username, email, password string) string {
	row := checkUser.QueryRow(username, password, email)
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
