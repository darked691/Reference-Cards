package main

import "fmt"

func main() {
    s := make([]string, 2)
	s[0] = "a"
	s[1] = "b"
	s = append(s, "c") 
	fmt.Println(s)
}
