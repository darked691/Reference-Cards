package main

import "fmt"

func main() {
    var b [2] int
	b[0]=0
	b[1]=1
	fmt.Println(b)

	/* ou directement */ 

	c := [2]int{0, 1}
	fmt.Println(c)
}

