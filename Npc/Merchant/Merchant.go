package Npc

import (
	"fmt"
	char "jeu/Character"
	item "jeu/item"
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
func ChoiceMerchant(perso *char.Character) {
	var choice string
	for {
		fmt.Println("\n1. Seringues de combat (+46% dgt/tour sur 2 tour) -> 55 or ")
		fmt.Println("\n2. Seringues de soin (+50PV instantanément) -> 45 or ")
		fmt.Println("\n3. Grenade incendiaires (-10%PV créatures/tour sur 3 tour) -> 40 or")
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
