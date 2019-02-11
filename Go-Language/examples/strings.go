package main

import "fmt"
import "strings"

func main() {
	// Count compte le nombre de sous-chaines ("ca") dans un string ("cacao")
	nb1 := strings.Count("cacao", "ca")
	fmt.Println("Il y a", nb1, "\"ca\" dans \"cacao\"")

	// Index renvoie le numéro de la première case de la sous-chaine ("a") dans un string ("cacao")
	nb2 := strings.Index("cacao", "a")
	fmt.Println(nb2)

	// Join concatène un ensemble de string par un séparateur 
	// (si on ne souhaite pas de séparateur il faut mettre "")
	chaine1 := strings.Join([]string{"clement", "mehdi", "TER", "master"}, ", ")
	fmt.Println(chaine1)

	// Repeat concatène un certain nombre de fois (4) un string ("ah")
	chaine2 := strings.Repeat("ah",4)
	fmt.Println(chaine2)

}