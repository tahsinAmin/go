// package main

// import "fmt"

// // func updateName(n string) {
// // 	n = "wedge"
// // }

// func updateName(n *string) {
// 	*n = "wedge"
// }

// func main() {
// 	name := "tofy"

// 	// updateName(name)
// 	// fmt.Println(name)``

// 	// & gets the memory address of the value (pointer)
// 	fmt.Println("memory address of name is:", &name)

// 	// * gets the value at the specified memory address
// 	m := &name
// 	fmt.Println("memory address:", m)
// 	fmt.Println("value at memory address:", *m)

// 	updateName(m) // using pointer as arg, can pass &name directly instead of "m" as well
// 	fmt.Println(name)
// }
