package main

import (
	"fmt"
	"math"
)

func sayMorning(n string) {
	fmt.Printf("Good Morning %v\n", n)
}
func sayNight(n string) {
	fmt.Printf("Good Night %v\n", n)
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func main() {
	fmt.Printf("%0.2f\n", circleArea(2))
	fmt.Printf("%0.2f\n", circleArea(2.5))

	cycleNames([]string{"mario", "luigi", "yoshi"}, sayMorning)
}
