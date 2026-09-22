package jeu

import (
	"fmt"
	jeu "jeu/Character"
	"time"
)

func Merchant() {

	texte := "\nBonjour jeune aventurier ! Bienvenur dans mon magasin, qu'est-ce qu'il te ferais plaisir ?"

	for _, lettre := range texte {
		fmt.Printf("%c", lettre)
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println()

}
func Shop(perso *jeu.Character) {
	var choice string
	for {
		fmt.Println("\n1. Seringues de combat (+46% dgt/tour sur 2tour) -> 55 or")
		fmt.Println("\n2. Seringues de soin (+50PV intstantanément) -> 45 or ")
		fmt.Println("\n3. Grenade incendiaires (-10%PV créatures/tour sur 3 tour)")
		fmt.Println("\n4. Quitter")
		fmt.Print("Ton choix (1/2/3/4) : ")
		fmt.Scan(&choice)

		switch choice {
		case "1":
			perso.buy(item.NewCombatSyringe(1))
		case "2":
			perso.buy(item.NewHealSyringe(1))
		case "3":
			perso.buy(item.NewGrenadeIncendiary(1))
		case "4":
			fmt.Println("À bientôt !")
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
