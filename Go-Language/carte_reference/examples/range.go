package main

import "fmt"

func main() {
    tab := []int{1, 2, 3}
	somme := 0
	for _, val := range tab {
    	somme += val
	}
	fmt.Println("Somme:", somme)
}