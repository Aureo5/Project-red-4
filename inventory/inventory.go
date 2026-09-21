package jeu

import (
	"fmt"
	"jeu/Project-red-4/menu"
)

func NewInv() {

}

func Inv() {
	if inv == 0 {
		fmt.Println("Il n'y a rien dans l'inventaire")
		jeu.Menu()
	}
}
