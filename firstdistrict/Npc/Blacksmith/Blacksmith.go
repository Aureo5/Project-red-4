package jeu

import (
	"fmt"
	"time"
)

func Blacksmith() {
	texte := "\nBonjour jeune aventurier ! Bienvenue dans mon atelier, que veux-tu forger ?"
	for _, lettre := range texte {
		fmt.Printf("%c", lettre)
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println()

}
func ChoiceBlacksmith(perso) {
	var choice string
	for {
		fmt.Println("\n1. Propulseur améliorer : Augmentation du réservoir (+10% gaz); Augmentation des PV (+10PV); Augmentation de l'initiative (+7%) -> 2 os, 3 peau, 50 or")
		fmt.Println("\n2. Grappin améliorer : Augmentation de la solidité du cable (+4PV); Augmentation de la stabilité du cable (+2 dgt) -> 3 os, 1 peau, 40 or")
		fmt.Println("\n3. Lame améliorer : Augmentation des dégats (+10 dgt); Augmentation de la durabilité (+50%) -> 5 os, 2 peau, 70 or")
		fmt.Println("\n4. Quitter")
		fmt.Print("Ton choix (1/2/3/4) : ")
		fmt.Scan(&choice)

		switch choice {
		case "1":
			perso.Buy(item.NewCombatSyringe(1))
		case "2":
			perso.Buy(item.NewHealSyringe(1))
		case "3":
			perso.Buy(item.NewGrenadeIncendiary(1))
		case "4":
			fmt.Println("À bientôt !")
			return
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
