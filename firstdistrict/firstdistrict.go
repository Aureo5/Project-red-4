package jeu

import (
	"fmt"
	"time"
)

func Firstdistrict() {

	texte := "Bienvenue dans le district d'Heraclion ! \nLes géants ont envahi ce district... Libérez les murs de ces abominations et accédez au prochain district !"

	for _, lettre := range texte {
		fmt.Printf("%c", lettre)
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println()

	Choicefirstdistrict()
}

func Choicefirstdistrict() {
	var choix string

	for {
		fmt.Println("\n1. Allez au camp")
		fmt.Println("2. Combattre les géants")
		fmt.Println("3. Accéder au prochain district")
		fmt.Print("Ton choix : ")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			fmt.Println("Bienvenue au camp ! Ici tu pourras acheter des objets et sauvegarder ta partie !")
			return
		case "2":
			fmt.Println("Bienvenue sur le champ de bataille ! Sois prudent, les géants peuvent être plus dangereux que tu ne le penses...")
			return
		case "3":
			fmt.Println("Attention ! Un titan gueule de loup protège le prochain mur, affronte le avant d'accéder au prochain district !")
			return
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
