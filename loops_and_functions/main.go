package main

import (
	"fmt"
	"math"
)

const delta = 1e-8

func Sqrt(x float64) float64 {

	z := x

	for {

		new_z := z - (z*z-x)/(2*z)

		if math.Abs(z-new_z) < delta {
			return z
		}

		z = new_z

	}

}

func main() {
	fmt.Println(Sqrt(2))
}
