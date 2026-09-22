package jeu

import (
	"fmt"
	CreateCharacter "jeu/character"
	"time"
)

func Lore() {
	texte := "\nSur un continent lointain...l'humanité ce fait assaillir par des immondisses aux tailles anormaux.\nLes habitants ce réfugis au centre du royaume d'olympe, dèrrière 3 murs gargantuesque: Le mur Midas qui entoure le district d'Heraclion; le mur Hermes \nqui encercle le district d'Argos et enfin le mur Arès qui ceinture le district d'Athènes ravagé par une invasion de créatures.\n"
	for _, lettre := range texte {
		fmt.Printf("%c", lettre)
		time.Sleep(20 * time.Millisecond)
	}
	CreateCharacter.CreateCharacter()
}
