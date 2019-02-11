package main

import "fmt"
import "os"

func main() {
    f, err1 := os.OpenFile("test_ouverture.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err1 != nil {
		panic(err1)
	}
	nb_octets, err2 := f.WriteString("voici un test")
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("Nb d'octets écrits dans le fichier : ", nb_octets)
	err3 := f.Close();
	if err3 != nil {
		panic(err3)
	}
}