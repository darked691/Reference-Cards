package main

import "fmt"

func main() {
    s := make([]string, 2)
	s[0] = "a"
	s[1] = "b"
	fmt.Println(s)

	/* ou directement */ 

	t := []string{"a", "b"}
	fmt.Println(t)
}
