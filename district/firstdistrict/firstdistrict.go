package district

import (
	"fmt"
	"jeu/Camp"
	char "jeu/Character"
	inv "jeu/Character/inventory/inventoryInteractive"
	ennemy "jeu/Npc/Ennemy"
	second "jeu/district/seconddistrict"
	fight "jeu/fight"
	"time"
)

func StartingPoint(perso *char.Character) {
	texteun := "\nUn titan colossale est apparue derrière le mur Arès et a créé une ouverture pour ses confrères !\nIl est de votre devoir de sauver les habitants tout en éliminant ces immondices. \nEn avant soldat !\n"
	for _, lettre := range texteun {
		fmt.Printf("%c", lettre)
		time.Sleep(20 * time.Millisecond)
	}

	textedeux := "\n ====Mission 1==== \nBienvenue dans le district d'Heraclion ! \nLes géants ont envahi ce district... Libérez les murs de ces abominations et accédez au prochain district !"

	for _, lettre := range textedeux {
		fmt.Printf("%c", lettre)
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println()

	Choicefirstdistrict(perso)
}

func Choicefirstdistrict(perso *char.Character) {
	var choix string

	for {
		fmt.Println("\n1. Allez au camp")
		fmt.Println("2. Combattre les géants")
		fmt.Println("3. Accéder au prochain district")
		fmt.Println("4. Inventaire")
		fmt.Print("Ton choix : ")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			fmt.Println("\nBienvenue au camp ! Ici tu pourras acheter des objets et sauvegarder ta partie !")
			jeu.Choicecamp(perso)
		case "2":
			fmt.Println("Bienvenue sur le champ de bataille ! Sois prudent, les géants peuvent être plus dangereux que tu ne le penses...")
			fight.StartFight(perso, &ennemy.Giant)

		case "3":
			fmt.Println("Attention ! Un titan gueule de loup protège le prochain mur, affronte le avant d'accéder au prochain district !")
			fight.BossFight(perso, &ennemy.Wolftitan)
			if perso.Unlock {
				perso.Unlock = false
				second.SecondPoint(perso)
			}

		case "4":
			inv.DisplayCharacterMenu(perso)

		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
