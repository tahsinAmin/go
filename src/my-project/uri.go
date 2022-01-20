package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	Println := fmt.Println
	var (
		lines []string
		a     int
		b     int
	)
	fmt.Scan(&a, &b)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	lines = strings.Split(text, " ")
	Println(lines)
}

// https://www.beecrowd.com.br/judge/en/problems/view/1196
