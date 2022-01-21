export GO111MODULE=on

# tutorial following - https://www.youtube.com/watch?v=ypV7r1ODZCA&list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM&index=10

# Installation of Go

- https://www.cyberciti.biz/faq/how-to-install-gol-ang-on-ubuntu-linux/#testing

# What I've leanred

- You can't initialize a string variable with a single quote
- You cannot use the short hand version of initializing outside the function
- You cannot run a file if a variable is not used.
- float 64 gives a better precision.
- A project block will have only one main method.
- A library gets removed autromatrically if its of no use anymore. similarly, a library is import if its being used
- ```
  Println := fmt.Println
  Println("Hello")
  ```

- The fields in the struct has to be capitalize

# Problems and it solutions

- [x] Go: "stat hello.go: no such file or directory". Solution is to see if both the files are in the same directory.
- [ ] .Sprintf("%-25v\n", value). But what about .Sprintf("%0.2f\n", value)?
- [x] Getting into that bin

```
nano ~/.profile
```

```
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:/home/w3e65/go/bin

export GOPATH="/home/w3e65/go/"
export GOBIN="/home/w3e65/go/bin"
```

- How to use EOF?

```

```
