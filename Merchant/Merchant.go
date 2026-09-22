package jeu

import (
	"fmt"
	"time"
)

func Merchant() {

	texte := "Bonjour jeune aventurier ! Bienvenur dans mon magasin, qu'est-ce qu'il te ferais plaisir ?"

	for _, lettre := range texte {
		fmt.Printf("%c", lettre)
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println()

}
