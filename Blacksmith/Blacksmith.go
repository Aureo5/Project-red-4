package jeu

import (
	"fmt"
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
