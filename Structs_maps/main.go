package main

import "fmt"

func main() {

	fam := map[string]string{
		"Guillermo": "Papa",
		"Elena":     "Mama",
		"Valeria":   "Hija",
		"Otro":      "Otro",
	}

	fmt.Println(fam["Guillermo"])
	delete(fam, "Otro")
	fmt.Println(fam)
	pop, ok := fam["Guillermo"]
	fmt.Println(pop, ok)
}
