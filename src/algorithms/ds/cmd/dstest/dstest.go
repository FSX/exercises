package main

import (
	"bufio"
	"image/png"
	"os"

	"algorithms/ds"
)

func main() {
	m, err := ds.NewMap(12.0, 5.5, 24.82, 42.2, 122.9)
	if err != nil {
		panic(err)
	}

	m.Generate(0.3)

	fd, err := os.Create("test.png")
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(fd)

	err = png.Encode(w, m.Image())
	if err != nil {
		panic(err)
	}

	w.Flush()
	fd.Close()
}
