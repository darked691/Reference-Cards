package main

import "fmt"
import "errors"

func verifie_positif (nombre float32){
    if(nombre < 0){
        panic(errors.New("Ce nombre ne doit pas etre negatif"))
    }
}

func quotient (num float32, denom float32) (float32, error){
    if denom == 0 {
        return -1, errors.New("Erreur mathematique")
    }else {
        return num / denom, nil
    }
}

func main() {
    var c1 float32
    var err error
    var a1, b1 float32 = 1,2
    c1,err = quotient(a1, b1)
    fmt.Println(c1)
    fmt.Println(err)
    verifie_positif(12.0)
}