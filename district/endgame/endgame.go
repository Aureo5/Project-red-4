package district

import (
	"fmt"
	chara "jeu/Character"
	"os"
	"time"
)

func Endgame(perso *chara.Character) {
	texte := "\n ====Fin du jeu==== \nMerci jeune aventurier, grâce à toi nous avons récupérer nos terres et éliminé toutes les créatures !"

	for _, lettre := range texte {
		fmt.Printf("%c", lettre)
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println()
}
func Choiceendgame(perso *chara.Character) {
	var choix string

	for {
		fmt.Println("\n1. Qui sont-ils ?")
		fmt.Println("2. Quitter le jeu")
		fmt.Print("Ton choix : ")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			fmt.Println("\n Arthur Commanay & Aurélien Heng - Bachelor 1, Cybersécurité, Ynov campus val d'europe")
			os.Exit(0)
		case "2":
			os.Exit(0)
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
