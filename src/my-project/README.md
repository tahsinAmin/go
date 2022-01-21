https://yourbasic.org/golang/split-string-into-slice/

# Things I've learned

- How to take a string and split to a list:

```
reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	lines = strings.Split(text, " ")
```

-     Make an array after knowing the length:

```
fmt.Scan(&no_of_people, &questions)

	var problems_solved [4]int
	Println(problems_solved[0])
```
