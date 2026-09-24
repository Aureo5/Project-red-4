package Npc

import (
	"fmt"
	char "jeu/Character"
	equip "jeu/equipment"
	"time"
)

func Lastblacksmith() {
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
		fmt.Println("\n1. Propulseur d'Ursadon : Augmentation des PV (+40PV) -> Peau d'Ursadon - 5 peau - 45 or ")
		fmt.Println("\n2. Grappin d'Ursadon : Augmentation de la solidité du câble (+15PV); Augmentation de la stabilité du câble (+15dgt) -> Peau d'Ursadon - 3 os - 3 peau - 30 or")
		fmt.Println("\n3. Lame d'Ursadon : Augmentation des dégâts (+40dgt) -> Peau d'Ursadon - 5 os - 60 or")
		fmt.Println("\n4. Quitter")
		fmt.Print("\nTon choix (1/2/3/4) : ")
		fmt.Scan(&choice)

		switch choice {
		case "1":
			perso.UpgradeSpecial("Propulsor", "bear_skin", "peau d'ours", 1, 0, 5, 45, func(e *equip.Equipment) {
				e.Healthbonus += 40
			})
		case "2":
			perso.UpgradeSpecial("Grapplin", "bear_skin", "peau d'ours", 1, 3, 3, 30, func(e *equip.Equipment) {
				e.Healthbonus += 15
				e.Damagebonus += 15
			})
		case "3":
			perso.UpgradeSpecial("Lame", "bear_skin", "peau d'ours", 1, 5, 0, 60, func(e *equip.Equipment) {
				e.Damagebonus += 40
			})
		case "4":
			fmt.Println("À bientôt !")
			return
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
