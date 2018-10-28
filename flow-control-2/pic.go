package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var list = make([][]uint8, dy)
	
	for i := range list {
		list[i] = make([]uint8, dx)
	}
	
	for i :=range list {
		for j := range list[i] {
			list[i][j] = uint8(i * j)
		}
	}
	
	return list
}

func main() {
	pic.Show(Pic)
}