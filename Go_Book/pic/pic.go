package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	rows := make([][]uint8, dy)
	for i := range rows {
		rows[i] = make([]uint8, dx)
	}
	for i, col := range rows {
		for j := range col {
			// Calculate the "blue-scale" value for position.
			y, x := uint8(i), uint8(j)
			// y*y - x*x https://imgur.com/rbE3DcJ
			// y*y + x*x https://imgur.com/L2TepyN
			col[j] = y*y - x*x
		}
	}
	return rows
}

func main() {
	pic.Show(Pic)
}

/*
	s\IMAGE:\data:image/png;base64,\
	IMAGE:iVBORw0KGgoAAAANSUhEUgAAAQAAAAEACA...
	Then use like <img src="..."></...
*/
