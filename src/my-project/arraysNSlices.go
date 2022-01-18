// package main

// import "fmt"

// func main() {
// 	var ages [3]int = [3]int{2, 3, 5}
// 	fmt.Println(ages)
// 	ages2 := [3]int{2, 3, 54}
// 	fmt.Println(ages2)
// 	var ages3 = [3]int{4, 34, 2}
// 	fmt.Println(ages3, len((ages3)))

// 	// slices (use arrays uynder the hood)
// 	scores := []int{23, 45, 124, 21}
// 	scores[2] = 21
// 	scores2 := append(scores, 85) // returns a new slice
// 	fmt.Println(scores2, len((scores2)))

// 	// slice range
// 	rangeOne := scores2[1:3]
// 	// fmt.Println(rangeOne)
// 	rangeTwo := scores2[1:]
// 	rangeThree := scores2[:3]

// 	rangeThree = append(rangeThree, 100)
// 	fmt.Println(rangeOne, rangeTwo, rangeThree)
// }
