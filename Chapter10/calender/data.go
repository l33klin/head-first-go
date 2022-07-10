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

type NewData struct {
	year int
}

func (n *NewData) SetYear(y int) {
	n.year = y
}

func (n *NewData) Year() int {
	return n.year
}
