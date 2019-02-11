package main

import "fmt"

func main() {
    var a int = 3
	if a < 0{
		fmt.Println(a, "negatif")
	} else if a > 0 {
		fmt.Println(a, "positif")
	} else {
		fmt.Println(a, "nul")
	}
}

