package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	addUser         *sql.Stmt
	getUserByToken  *sql.Stmt
	checkRepeatUser *sql.Stmt
	checkUser       *sql.Stmt
	selectByDay     *sql.Stmt
	selectByMonth   *sql.Stmt
	addEvent        *sql.Stmt
	delEvent        *sql.Stmt
)

func Init() error {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/calender?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	//users
	addUser, err = db.Prepare("INSERT INTO users(`username`, `email`, `password`, `token`) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	getUserByToken, err = db.Prepare("SELECT `id` FROM users WHERE `token` = ?")
	if err != nil {
		return err
	}
	checkRepeatUser, err = db.Prepare("SELECT `id` FROM users WHERE email=? OR username=?")
	if err != nil {
		return err
	}
	checkUser, err = db.Prepare("SELECT `token` from users WHERE `username`=? AND `password`=? AND `email`=?")
	if err != nil {
		return err
	}
	//event
	selectByDay, err = db.Prepare("SELECT `id`, `title`, `time` FROM events WHERE year=? AND month=? AND day=? AND uid=?")
	if err != nil {
		return err
	}
	selectByMonth, err = db.Prepare("SELECT DISTINCT `month` FROM events WHERE year=? AND month=? AND uid=?")
	if err != nil {
		return err
	}
	addEvent, err = db.Prepare("INSERT INTO events(`uid`, `year`, `month`, `day`, `time`, `title`) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	delEvent, err = db.Prepare("DELETE FROM events WHERE `id`=? AND uid=?")
	if err != nil {
		return err
	}

	return nil
}

type EventModel struct {
	EventID string `json:"event_id"`
	Year    int    `json:"year"`
	Month   int    `json:"month"`
	Day     int    `json:"day"`
	Title   string `json:"title"`
	Time    string `json:"time"`
}
