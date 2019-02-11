package main

import "fmt"

func main() {
    a := 1
	switch a%2 {
		case 0:
			fmt.Println("pair")
		case 1:
			fmt.Println("impair")
	}
}
