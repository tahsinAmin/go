package main

import (
	"fmt"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// Println := fmt.Println
	var (
	// lines             []string
	// no_of_people      int
	// questions         int
	// no_one_solved_all bool
	)

	result := make([][]string, 4)
	fmt.Printf("%q, %q, %q, %q \n", result[0][0], result[0][1], result[0][2], result[0][3])

	fmt.Scan(&result[0][0], &result[0][1])

	// // problems_solved := make([]int, questions)
	// all := make([][]string, questions)
	// cnt_solved := 0
	// score := 0
	// // Println(problems_solved[0])

	// for x := 0; x < no_of_people; x++ {
	// 	fmt.Scan(&all[x][0], &all[x][1], &all[x][2])

	// 	// text, _ := reader.ReadString('\n')
	// 	// text = strings.Replace(text, "\n", "", -1)
	// 	// no_one_solved_all = strings.Contains(text, "0")
	// 	// lines = strings.Split(text, " ")
	// 	// all = append(all, lines)
	// 	// for x := 0; x < len(lines); x++ {
	// 	// 	if lines[x] == "1" {
	// 	// 		problems_solved[x] = 1
	// 	// 		cnt_solved += 1
	// 	// 	}
	// 	// 	if cnt_solved == questions {
	// 	// 		break
	// 	// 	}
	// 	// }
	// }

	// if no_one_solved_all {
	// 	score += 1
	// }
	// if cnt_solved == questions {
	// 	score += 1
	// }

	// Println(score)
	// Println(all)
}

// https://www.beecrowd.com.br/judge/en/problems/view/1514
