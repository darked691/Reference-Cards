package main

import "fmt"

type etudiant struct {
    nom string
    num  int
}

func main() {
    e := etudiant{nom: "Clement", num: 21501810}
    fmt.Println(e)
}