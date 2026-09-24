package Npc

import (
	"fmt"
	char "jeu/Character"
	equip "jeu/equipment"
	"time"
)

func Secblacksmith() {
	texte := "\nBonjour jeune aventurier ! Bienvenue dans mon atelier, que veux-tu forger ?"

	for _, lettre := range texte {
		fmt.Printf("%c", lettre)
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println()
}

func ChoiceSecblacksmith(perso *char.Character) {
	var choice string
	for {
		fmt.Println("\n1. Propulseur de loup : Augmentation des PV (+25PV) -> Peau de loup - 5 peau - 45 or ")
		fmt.Println("\n2. Grappin de loup : Augmentation de la solidité du câble (+10PV); Augmentation de la stabilité du câble (+10dgt) -> Peau de loup - 3 os - 3 peau - 30 or")
		fmt.Println("\n3. Lame de loup : Augmentation des dégâts (+20dgt) -> Peau de loup - 5 os - 60 or")
		fmt.Println("\n4. Quitter")
		fmt.Print("\nTon choix (1/2/3/4) : ")
		fmt.Scan(&choice)

		switch choice {
		case "1":
			perso.UpgradeSpecial("Propulsor", "wolf_skin", "peau de loup", 1, 0, 5, 45, func(e *equip.Equipment) {
				e.Healthbonus += 25
			})
		case "2":
			perso.UpgradeSpecial("Grapplin", "wolf_skin", "peau de loup", 1, 3, 3, 30, func(e *equip.Equipment) {
				e.Healthbonus += 10
				e.Damagebonus += 10
			})
		case "3":
			perso.UpgradeSpecial("Lame", "wolf_skin", "peau de loup", 1, 5, 0, 60, func(e *equip.Equipment) {
				e.Damagebonus += 20
			})
		case "4":
			fmt.Println("À bientôt !")
			return
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
