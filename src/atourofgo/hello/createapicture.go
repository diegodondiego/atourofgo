package main

import (
	"code.google.com/p/go-tour/pic"
)

func Pic(dx, dy int) (result [][]uint8) {
	result = make([][]uint8, dy)
	for pos, _ := range result {
		temp := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			temp[x] = uint8(x^pos)
		}
		result[pos] = temp
	}
	return
}

func main() {
	pic.Show(Pic)
}

