package jeu

import (
	"fmt"
	character "jeu/Project-red-4/character"
)

func DisplayInv(c character.Character) {
	if len(c.Inventaire) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}
	fmt.Println("--- INVENTAIRE ---")
	for _, item := range c.Inventaire {
		fmt.Printf("- %s (x%d) : %.2f gold\n", item.Nom, item.Quantity, item.Prix)
	}
}
