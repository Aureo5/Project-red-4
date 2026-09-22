package jeu

import (
	"Jeu/firstshop"
	"fmt"
)

func Choicecamp() {
	var choix string

	for {
		fmt.Println("\n===que veux-tu faire?=== ")
		fmt.Println("\n1. Voir le marchand")
		fmt.Println("2. Voir le forgeron")
		fmt.Println("3. Sauvegarder la partie")
		fmt.Print("Ton choix : ")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			Merchant()
			return
		case "2":
			fmt.Println
			return
		case "3":
			fmt.Println
			return
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
