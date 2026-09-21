package jeu

import (
	"fmt"
	"time"
)

func firstdistrict() {

	texte := "Bienvenue dans le district d'Heraclion ! \nLes géants ont envahi ce district... Libérez les murs de ces abominations et accédez au prochain district !"

	for _, lettre := range texte {
		fmt.Printf("%c", lettre)
		time.Sleep(20 * time.Millisecond)
	}
}
