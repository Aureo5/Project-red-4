package jeu

import (
	"fmt"
	char "jeu/Character"
	ennemy "jeu/Npc/Ennemy"
	Blacksmith "jeu/Npc/Lastblacksmith"
	Merchant "jeu/Npc/Merchant"
	training "jeu/fight/training"
	"os"
)

func Choicelastcamp(perso *char.Character) {
	var choix string

	for {
		fmt.Println("\n===que veux-tu faire?=== ")
		fmt.Println("\n1. Voir le marchand")
		fmt.Println("2. Voir le forgeron")
		fmt.Println("3.Entrainement")
		fmt.Println("4.Revenir au district")
		fmt.Println("5.Quitter le jeu")
		fmt.Print("Ton choix : ")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			Merchant.Merchant()
			Merchant.ChoiceMerchant(perso)
			return
		case "2":
			Blacksmith.Lastblacksmith()
			Blacksmith.ChoiceSecblacksmith(perso)
			return
		case "3":
			training.TrainingFight(perso, &ennemy.Coloss)
		case "4":
			return
		case "5":
			os.Exit(0)
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
