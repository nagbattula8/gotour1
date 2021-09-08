package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

const delta = 1e-8

func (e ErrNegativeSqrt) Error() string {

	return fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := x

	for {
		new_z := z - (z*z-x)/(2*z)

		if math.Abs(new_z-z) < delta {
			return new_z, nil
		}

		z = new_z
	}

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
