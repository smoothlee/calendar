package mysql

func Init() {

}

type EventModel struct {
	EventID string `json:"event_id"`
	Year    int    `json:"year"`
	Month   int    `json:"month"`
	Day     int    `json:"day"`
	Title   string `json:"title"`
	Text    string `json:"text"`
}
