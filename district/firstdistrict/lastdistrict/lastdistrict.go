package district

import (
	"fmt"
	jeu "jeu/Camp"
	chara "jeu/Character"
	ennemy "jeu/Npc/Ennemy"
	"jeu/fight"
	"time"
)

func SecondPoint(perso *chara.Character) {
	texte := "\n ====Mission 3==== \nBienvenue dans le district d'Athènes ! \nMerci d'avoir libérez les murs Hermes ! Les colosses ont envahi ce district mais ils protègent leur maître, éliminés les afin de libérer le pays d'Osnos de ce cauchemar !"

	for _, lettre := range texte {
		fmt.Printf("%c", lettre)
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println()

}
func Choicelastdistrict(perso *chara.Character) {
	var choix string

	for {
		fmt.Println("\n1. Allez au camp")
		fmt.Println("2. Combattre les colosses")
		fmt.Println("3. Combattre le dernier boss")
		fmt.Println("4. Inventaire")
		fmt.Print("Ton choix : ")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			fmt.Println("\nBienvenue au camp ! Ici tu pourras acheter des objets et sauvegarder ta partie !")
			jeu.Choicelastcamp(perso)
		case "2":
			fmt.Println("Bienvenue sur le champ de bataille ! Sois prudent, les géants peuvent être plus dangereux que tu ne le penses...")
			fight.StartFight(perso, &ennemy.Coloss)
		case "3":
			fmt.Println("Attention ! Un titan gueule de loup protège le prochain mur, affronte le avant d'accéder au prochain district !")
			fight.StartFight(perso, &ennemy.Colossaletitan)
		case "4":
			perso.DisplayCharacterMenu()
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
