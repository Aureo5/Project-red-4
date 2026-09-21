package jeu

import (
	"fmt"
	"time"
)

func Lore() {
	texte := "Sur un continent lointain...l'humanité ce fait assaillir par des immondisses aux tailles anormaux. Les habitants ce réfugis au centre du royaume d'olympe, dèrrière 3 murs gargantuesque: Le mur Midas qui entoure le district d'Heraclion; le mur Hermes qui encercle le district d'Argos et enfin le mur Arès qui ceinture le district d'Athènes ravagé par une invasion de créatures."
	for _, lettre := range texte {
		fmt.Printf("%c", lettre)
		time.Sleep(50 * time.Millisecond)
	}
	jeu.CreateCharacter()
}

func StartingPoint() {

}
