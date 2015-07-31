// A Go port from http://www.playfuljs.com/realistic-terrain-in-130-lines/
// Found here: http://blog.habrador.com/2013/02/how-to-generate-random-terrain.html
// And: http://www.gameprogrammer.com/fractal.html
// This doesn't give the expected result. I am an idiot and I don't know
// how to fix it.

package ds

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"

	"github.com/mgutz/ansi"
)

var ErrMapSize = errors.New("map size not power of two plus one")

type Map struct {
	size  int
	max   int
	array []float64
}

func NewMap(size, nw, ne, sw, se float64) (*Map, error) {
	s := int(math.Pow(2.0, size)) + 1

	a := make([]float64, s*s)
	m := &Map{s, s - 1, a}

	rand.Seed(345)

	// Set corner values.
	a[0] = nw
	a[s-1] = ne
	a[s*(s-1)] = sw
	a[s*s-1] = se

	return m, nil
}

func (m *Map) get(x, y int) float64 {
	if x < 0 || y < 0 || x > m.max || y > m.max {
		return 0.00
	}

	return m.array[x+y*m.size]
}

func (m *Map) set(x, y int, v float64) {
	m.array[x+y*m.size] = v
}

func (m *Map) Generate(roughness float64) {
	m.divide(m.max, roughness)
}

func (m *Map) divide(size int, roughness float64) {
	h := size / 2
	s := roughness * float64(size)

	if h < 1 {
		return
	}

	for y := h; y < m.max; y += size {
		for x := h; x < m.max; x += size {
			m.square(x, y, h, rand.Float64()*s*2-s)
		}
	}

	for y := 0; y <= m.max; y += h {
		for x := (y + h) % size; x <= m.max; x += size {
			m.diamond(x, y, h, rand.Float64()*s*2-s)
		}
	}

	m.divide(size/2, roughness)
}

func (m *Map) square(x, y, size int, offset float64) {
	avg := (m.get(x-size, y-size) + // Upper left
		m.get(x+size, y-size) + // Upper right
		m.get(x+size, y+size) + // Lower right
		m.get(x-size, y+size)) / 4 // Lower left

	m.set(x, y, avg+offset)
}

func (m *Map) diamond(x, y, size int, offset float64) {
	avg := (m.get(x, y-size) + // Upper left
		m.get(x+size, y) + // Upper right
		m.get(x, y+size) + // Lower right
		m.get(x-size, y)) / 4 // Lower left

	m.set(x, y, avg+offset)
}

func (m *Map) String() string {
	var b bytes.Buffer

	// Nice colors for debugging.
	r := ansi.ColorFunc("red")
	g := ansi.ColorFunc("green")

	for y := 0; y < m.size; y++ {
		for x := 0; x < m.size; x++ {
			v := m.get(x, y)

			if v > 0 {
				fmt.Fprint(&b, r(fmt.Sprintf("%.2f ", v)))
			} else {
				fmt.Fprint(&b, g(fmt.Sprintf("%.2f ", v)))
			}
		}

		fmt.Fprint(&b, "\n")
	}

	return b.String()
}

func (m *Map) Image() image.Image {
	i := image.NewGray(image.Rect(0, 0, m.size, m.size))

	for y := 0; y < m.size; y++ {
		for x := 0; x < m.size; x++ {
			v := m.get(x, y)

			// This gives a nice result.
			i.SetGray(x, y, color.Gray{uint8(v * (256 / 10))})
		}
	}

	return i
}
