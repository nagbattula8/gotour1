package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	dp := [][]uint8{}

	for i := 0; i < dy; i++ {

		dp = append(dp, make([]uint8, dx))
	}

	return dp

}

func main() {
	pic.Show(Pic)
}
