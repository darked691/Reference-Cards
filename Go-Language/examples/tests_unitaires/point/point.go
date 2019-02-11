// Package point concernant tout ce qui concerne l'entité Point. 
package point

import "strconv"

// Point est une structure ayant une abscisse et une ordonnée. 
type Point struct {
	// X abscisse du point
	X int

	// Y ordonnee du point
	Y int
}

// PrintPoint affiche l'abscisse suivie de l'ordonnée. 
func PrintPoint (p Point) string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}

// PrintPointReverse affiche l'ordonnée suivie de l'abscisse. 
func PrintPointReverse (p Point) string {
	return "(" + strconv.Itoa(p.Y) + "," + strconv.Itoa(p.X) + ")"
}


