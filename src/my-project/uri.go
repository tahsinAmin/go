package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	keys := map[string]string{
		"1":  "`",
		"2":  "1",
		"3":  "2",
		"4":  "3",
		"5":  "4",
		"6":  "5",
		"7":  "6",
		"8":  "7",
		"9":  "8",
		"0":  "9",
		"-":  "0",
		"=":  "-",
		"W":  "Q",
		"E":  "W",
		"R":  "E",
		"T":  "R",
		"Y":  "T",
		"U":  "Y",
		"I":  "U",
		"O":  "I",
		"P":  "O",
		"[":  "P",
		"]":  "[",
		"\\": "]",
		"S":  "A",
		"D":  "S",
		"F":  "D",
		"G":  "F",
		"H":  "G",
		"J":  "H",
		"K":  "J",
		"L":  "K",
		";":  "L",
		"'":  ";",
		"X":  "Z",
		"C":  "X",
		"V":  "C",
		"B":  "V",
		"N":  "B",
		"M":  "N",
		",":  "M",
		".":  ",",
		"/":  ".",
		" ":  " ",
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		for i := 0; i < len(str); i++ {
			fmt.Print(keys[str[i:i+1]])
		}
		fmt.Println("")
	}
}

// https://www.beecrowd.com.br/judge/en/problems/view/1196
