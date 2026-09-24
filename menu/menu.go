package menu

import (
	"fmt"
	character "jeu/Character"
	inv "jeu/Character/inventory/inventorydisplay"
	lore "jeu/lore"
	"os"
)

func Menu() {
	choix := 0
	fmt.Println("1.Start")
	fmt.Println("2.Personnage")
	fmt.Println("3.Inventaire")
	fmt.Println("4.Quit")
	fmt.Printf("Ton choix: ")

	fmt.Scan(&choix)
	var perso character.Character

	switch choix {
	case 1:
		lore.Lore()
	case 2:
		if perso.Name == "" {
			fmt.Println("\nLancer la partie pour créer votre personnage")
			Menu()
		} else {
			inv.Displayinfo(perso)
		}
	case 3:
		inv.DisplayInv(perso)
		inv.DisplayEquip(perso)
		Menu()
	case 4:
		os.Exit(0)
	default:
		fmt.Println("Choix invalide")
	}
}
