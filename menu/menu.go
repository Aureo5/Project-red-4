package jeu

import (
	"fmt"
	character "jeu/Project-red-4/character"
	inv "jeu/Project-red-4/inventory"
	lore "jeu/Project-red-4/lore"
	"os"
)

func Menu() {
	choix := 0
	fmt.Println("1.Start")
	fmt.Println("2.Personnage")
	fmt.Println("3.Inventaire")
	fmt.Println("4.Quit")
	fmt.Printf("Que faire: ")

	fmt.Scan(&choix)
	var perso character.Character

	switch choix {
	case 1:
		lore.Lore()
	case 2:
		if perso.Name == "" {
			fmt.Println("Lancer la partie pour créer votre personnage")
			Menu()
		} else {
			perso.Info()
		}
	case 3:
		inv.DisplayInv(perso)
		Menu()
	case 4:
		os.Exit(0)
	default:
		fmt.Println("Choix invalide")
	}
}
