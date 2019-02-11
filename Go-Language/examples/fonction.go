package main

import "fmt"

func calcul1 (val1 int, val2 int) int {
    return val1 + val2
}

func calcul2 (val1 int, val2 int) (int, int){
    return val1 + val2, val1 - val2
}

func main() {
    var test1, test2, test3 int = 0, 0, 0
    test1, test2 = calcul2(2,3)
    fmt.Println(test1)
    fmt.Println(test2)
    test3 = calcul1(1,2)
    fmt.Println(test3)
}