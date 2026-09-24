package district

import (
	"fmt"
	"jeu/Camp"
	chara "jeu/Character"
	inv "jeu/Character/inventory/inventoryInteractive"
	ennemy "jeu/Npc/Ennemy"
	last "jeu/district/lastdistrict"
	"jeu/fight"
	"time"
)

func SecondPoint(perso *chara.Character) {
	texte := "\n ====Mission 2==== \nBienvenue dans le district d'Argos ! \nMerci d'avoir libérez les murs Midas ! Les colosses ont envahi ce district, nous comptons sur vous pour nous libérez de ces abominations et accédez au dernier district !"

	for _, lettre := range texte {
		fmt.Printf("%c", lettre)
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println()
	Choiceseconddistrict(perso)
}
func Choiceseconddistrict(perso *chara.Character) {
	var choix string

	for {
		fmt.Println("\n1. Allez au camp")
		fmt.Println("2. Combattre dans le district")
		fmt.Println("3. Accéder au prochain district")
		fmt.Println("4. Inventaire")
		fmt.Print("Ton choix : ")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			fmt.Println("\nBienvenue au camp ! Ici tu pourras acheter des objets.")
			jeu.Choiceseccamp(perso)
		case "2":
			fmt.Println("Des grondements de pas raisonnent de partout...")
			fight.LaunchRandomFight(perso)
		case "3":
			fmt.Println("Attention ! Un Géant Ursadon protège le prochain mur, affronte le avant d'accéder au prochain district !")
			fight.BossFight(perso, &ennemy.Beartitan)
			if perso.Unlock {
				perso.Unlock = false
				last.LastPoint(perso)
			}
		case "4":
			inv.DisplayCharacterMenu(perso)
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
