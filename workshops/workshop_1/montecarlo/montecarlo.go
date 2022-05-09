package montecarlo

import (
	"fmt"
	"math"
	"math/rand"
)

// Use the Monte Carlo simulation to approximate the value of pi.
// Input: 100000
// Output: 3.1518 (just an example, doesn’t mean that with 100000 points you always going to get this value)
// You can consider that your program is working if with higher number of points your pi value is
// near to the real value of pi : 3.14159

func EstimatePi() {
	fmt.Println("Challenge 3: Estimate PI with Montecarlo")
	n := 10000000
	sum := 0
	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()

		squares_sum := math.Pow(x, 2) + math.Pow(y, 2)
		z := math.Sqrt(squares_sum)
		if z < 1 {
			sum++
		}
	}

	area := 4 * float64(sum) / float64(n)
	fmt.Println("Pi = ", area)
}
