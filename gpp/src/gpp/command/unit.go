package command

import "fmt"

type Unit struct {
	Name string
	x, y int
}

func (u *Unit) Position() (int, int) {
	return u.x, u.y
}

func (u *Unit) MoveTo(x, y int) {
	u.x, u.y = x, y
	fmt.Printf("%s moved to (%d,%d)\n", u.Name, x, y)
}
