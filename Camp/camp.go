package jeu

import (
	"fmt"
	Blacksmith "jeu/Blacksmith"
	Merchant "jeu/Merchant"
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
			Merchant.Merchant()
			return
		case "2":
			Blacksmith.Blacksmith()
			return
		case "3":
			return
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
