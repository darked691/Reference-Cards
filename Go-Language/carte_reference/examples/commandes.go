package main

import "fmt"
import "os"

func main() {
	// os.Args récupére les différents mots constituant la commande
	commandeTotale := os.Args
	// os.Args[1:] récupère tous les mots sauf le premier (appel du programme)
    commandeSansAppelProgr := os.Args[1:]

    fmt.Println(commandeTotale)
    fmt.Println(commandeSansAppelProgr)
}