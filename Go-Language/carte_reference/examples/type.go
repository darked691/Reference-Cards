package main

import "fmt"

func main() {
    var a byte = 60 // 00111100 en binaire
    var b byte = 13 // 00001101 en binaire

    fmt.Println("a&b = ", a & b) // AND : 00001100
    fmt.Println("a|b = ", a | b) // OR : 00111101
    fmt.Println("a^b = ", a ^ b) // XOR : 00110001
    fmt.Println("a<<1 = ", a << 1) // décalage de 1 vers la gauche : 01111000
    fmt.Println("a>>1 = ", a >> 1) //décalage de 1 vers la droite : 00011110

}
