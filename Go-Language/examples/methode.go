package main

import "fmt"

type etudiant struct {
    nom string
    num  int
}

func (e etudiant) change_num (nouveau int) etudiant{
    e.num=nouveau
    return e
}

func main() {
    e := etudiant{nom: "Clement", num: 21501810}
    e = e.change_num(21500000)
    fmt.Println(e)
}