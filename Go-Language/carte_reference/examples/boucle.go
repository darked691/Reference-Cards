package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	/* ou bien */

	j := 0 
	for j < 10 {
		fmt.Println(j)
		j = j + 1
	}
}

