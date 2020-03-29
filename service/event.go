package service

import "github.com/smoothlee/calendar/mysql"

func AddEvent(uid int64, year, month, day int, time, title string) error {
	event := &mysql.EventModel{
		Year:  year,
		Month: month,
		Day:   day,
		Title: title,
		Time:  time,
	}
	return mysql.AddEvent(uid, event)
}

func DelEvent(uid int64, eventID string) error {
	return mysql.DelEvent(uid, eventID)
}

func GetDay(uid int64, year, month, day int) ([]*mysql.EventModel, error) {
	return mysql.GetDay(uid, year, month, day)
}

func GetMonth(uid int64, year, month int) ([]int, error) {
	return mysql.GetMonth(uid, year, month)
}
