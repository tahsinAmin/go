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
		lines        []string
		no_of_people int
		questions    int
		solvedAll    bool
	)

	fmt.Scan(&no_of_people, &questions)

	problems_solved := make([]int, questions)
	cnt_solved := 0

	Println(problems_solved[0])

	for x := 0; x < no_of_people; x++ {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		solvedAll = strings.Contains(text, "0")
		lines = strings.Split(text, " ")
		for x := 0; x < len(lines); x++ {
			if lines[cnt_solved] == "1" {
				problems_solved[cnt_solved] = 1
				cnt_solved += 1
			} else {
				break
			}
		}
	}

	score := 0
	if solvedAll {
		score += 1
	}
	Println(score, lines)
	Println(problems_solved)

}

// https://www.beecrowd.com.br/judge/en/problems/view/1514
