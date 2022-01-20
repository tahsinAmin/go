https://yourbasic.org/golang/split-string-into-slice/

- How to take a string and split to a list:

```
reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	lines = strings.Split(text, " ")
```
