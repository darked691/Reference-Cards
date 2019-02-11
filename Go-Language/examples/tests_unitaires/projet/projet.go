// Package main contient la fonction principale pour l'exemple 
// Compilation/Tests Unitaires/Documentation. 
package main

import "fmt"
import "point"

// main est la fonction Principale du programme.
func main() {
    p := point.Point{X : 3, Y: 2}
    fmt.Println(point.PrintPoint(p))
}