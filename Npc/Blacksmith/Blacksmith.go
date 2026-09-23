package Npc

import (
	"fmt"
	char "jeu/Character"
	equip "jeu/equipment"
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

func ChoiceBlacksmith(perso *char.Character) {
	var choice string
	for {
		fmt.Println("\n1. Propulseur améliorer : Augmentation du réservoir (+30% gaz); Augmentation de l'initiative (+10%); Agmentation des PV (+10PV) -> 3 peau - 1 os - 45 or ")
		fmt.Println("\n2. Grappin améliorer : Augmentation de la solidité du câble (+5PV); Augmentation de la stabilité du câble (+4dgt) -> 2 os - 2 peau - 30 or")
		fmt.Println("\n3. Lame améliorer : Augmentation des dégâts (+12dgt); Augmentation de la durabilité (+40%) -> 5 os - 2 peau - 60 or")
		fmt.Println("\n4. Quitter")
		fmt.Print("Ton choix (1/2/3/4) : ")
		fmt.Scan(&choice)

		switch choice {
		case "1":
			perso.Upgrade("Propulsor", 3, 1, 45, func(e *equip.Equipment) {
				e.Gaz += e.Gaz * 30 / 100
				e.Speedbonus += e.Speedbonus * 10 / 100
				e.Healthbonus += 10
			})
		case "2":
			perso.Upgrade("Grapplin", 2, 2, 30, func(e *equip.Equipment) {
				e.Healthbonus += 5
				e.Damagebonus += 4
			})
		case "3":
			perso.Upgrade("Lame", 5, 2, 60, func(e *equip.Equipment) {
				e.Damagebonus += 12
				e.Durability += e.Durability * 40 / 100
			})
		case "4":
			fmt.Println("À bientôt !")
			return
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
