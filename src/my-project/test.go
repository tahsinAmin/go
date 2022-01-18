package main

import (
	"fmt"
	"strings"
)

func getInitial(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string
	for _, v := range names {
		initials = append(initials, v[0:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"

}

func main() {

	x, y := getInitial("Tahsin")
	fmt.Printf("%v - %v\n", x, y)
	// fmt.Printf("%0.2f\n", circleArea(2.5))

	// cycleNames([]string{"mario", "luigi", "yoshi"}, sayMorning)
}
