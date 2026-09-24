package jeu

import (
	"fmt"
	char "jeu/Character"
	Blacksmith "jeu/Npc/Blacksmith"
	Merchant "jeu/Npc/Merchant"
	"os"
)

func Choicecamp(perso *char.Character) {
	var choix string

	for {
		fmt.Println("\n===que veux-tu faire?=== ")
		fmt.Println("\n1. Voir le marchand")
		fmt.Println("2. Voir le forgeron")
		fmt.Println("3. Sauvegarder la partie")
		fmt.Println("4.Quitter le jeu")
		fmt.Print("Ton choix : ")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			Merchant.Merchant()
			Merchant.ChoiceMerchant(perso)
			return
		case "2":
			Blacksmith.Blacksmith()
			Blacksmith.ChoiceBlacksmith(perso)
			return
		case "3":
			return
		case "4":
			os.Exit(0)
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
