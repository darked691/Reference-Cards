package main

import "io/ioutil"

func main(){

	donnees_copiees, err1 := ioutil.ReadFile("test1.txt")
	if err1 != nil {
		panic(err1)
	}
	donnees_a_recopier := []byte(string(donnees_copiees))
	err2 := ioutil.WriteFile("test1_copie.txt", donnees_a_recopier, 0644)
	if err2 != nil {
		panic(err2)
	}
}