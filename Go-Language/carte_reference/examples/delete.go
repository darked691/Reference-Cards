package main

import "fmt"

func main() {
    m := make(map[string]int)
	m["mehdi"] = 1
	m["clement"] = 2
	fmt.Println(m)
	delete(m, "clement")
	fmt.Println(m)
}
