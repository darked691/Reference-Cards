package main

import "fmt"

func main() {
    s := make([]string, 2)
	s[0] = "a"
	s[1] = "b"
	c := make([]string, len(s))
	copy(c, s) 
	fmt.Println(s)
	fmt.Println(c)
}
