package jeu

import (
	"fmt"
	"jeu/Project-red-4/lore"
	"os"
)

func Menu() {
	choix := 0
	fmt.Println("1.Start")
	fmt.Println("2.Personnage")
	fmt.Println("3.Inventaire")
	fmt.Println("4.Quit")
	fmt.Printf("Que faire: ")

	fmt.Scan(&choix)
	switch choix {
	case 1:
		jeu.Lore()
	case 2:
		fmt.Println()
	case 3:

	case 4:
		os.Exit(0)

	}
}
