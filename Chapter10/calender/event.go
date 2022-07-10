package calender

type Event struct {
	Date
	NewData
	event string
}

func (e *Event) New(name string, y int) {
	e.event = name
	//e.year = y // Ambiguous reference
	e.Date.year = y
	e.NewData.year = y
}
