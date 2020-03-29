package mysql

func GetMonth(uid int64, year, month int) ([]int, error) {
	rows, err := selectByMonth.Query(year, month, uid)
	if err != nil {
		return nil, err
	}
	ret := []int{}
	for rows.Next() {
		var m int
		if err := rows.Scan(&m); err != nil {
			continue
		}
		ret = append(ret, m)
	}
	return ret, nil
}

func GetDay(uid int64, year, month, day int) ([]*EventModel, error) {
	rows, err := selectByDay.Query(year, month, day, uid)
	if err != nil {
		return nil, err
	}
	ret := []*EventModel{}
	for rows.Next() {
		e := &EventModel{
			Year:  year,
			Month: month,
			Day:   day,
		}
		if err := rows.Scan(&e.EventID, &e.Title, &e.Time); err != nil {
			continue
		}
		ret = append(ret, e)
	}
	return ret, nil
}

func DelEvent(id string) error {
	_, err := delEvent.Exec(id)
	return err
}

func AddEvent(uid int64, event *EventModel) error {
	_, err := addEvent.Exec(uid, event.Year, event.Month, event.Day, event.Time, event.Title)
	return err
}
