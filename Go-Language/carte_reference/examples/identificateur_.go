package main

import "fmt"

func main() {
    m := make(map[string]int)
	m["mehdi"] = 1
	m["clement"] = 2
	_, test1 := m["clement"]
	_, test2 := m["inconnu"]
	if test1 == true {
		fmt.Println("clement existe")
	}else {
		fmt.Println("clement n'existe pas")
	}
	if test2 == true {
		fmt.Println("inconnu existe")
	}else {
		fmt.Println("inconnu n'existe pas")
	}
}