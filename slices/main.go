package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

var ij int = 1

func Pic(dx, dy int) [][]uint8 {

	dp := [][]uint8{}

	for i := 0; i < dy; i++ {

		dp = append(dp, make([]uint8, dx))
	}

	return dp

}

func main() {

	fmt.Println(ij)
	pic.Show(Pic)
}
