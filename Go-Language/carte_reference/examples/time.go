package main

import "fmt"
import "time"

func main() {
    t1 := time.Now()
    fmt.Print(t1.Day(), " ", t1.Month(), " ", t1.Year(), " - ")
    fmt.Println(t1.Hour(), ":", t1.Minute(), ":", t1.Second(), ":", t1.Nanosecond())

    t2 := time.Now()
    difference := t2.Sub(t1)
    fmt.Println(difference)
}