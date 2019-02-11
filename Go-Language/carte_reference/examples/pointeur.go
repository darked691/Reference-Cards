package main

import "fmt"

func add (val *int){
    *val = *val + 1
}

func main() {
    var cal int = 2
    fmt.Println(cal)
    add(&cal)
    fmt.Println(cal)
}