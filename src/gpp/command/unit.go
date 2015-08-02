package command

// import "fmt"

type Unit struct {
	Name string
	X, Y int
}

func (u *Unit) Position() (int, int) {
	return u.X, u.Y
}

func (u *Unit) MoveTo(x, y int) {
	u.X, u.Y = x, y
	// fmt.Printf("%s moved to (%d,%d)\n", u.Name, x, y)
}
