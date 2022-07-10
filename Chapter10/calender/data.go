package calender

type Date struct {
	year int
}

func (d *Date) SetYear(y int) {
	d.year = y
}

func (d *Date) Year() int {
	return d.year
}
