package calender

type Event struct {
	Date
	event string
}

func (e *Event) New(name string, y int) {
	e.event = name
	e.year = y
	e.Date.year = y
}
