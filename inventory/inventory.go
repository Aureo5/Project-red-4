package jeu

import (
	"fmt"
	character "jeu/character"
)

func DisplayInv(c character.Character) {
	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}
	fmt.Println("--- INVENTAIRE ---")
	for _, item := range c.Inventory {
		fmt.Printf("- %s (x%d) : %.2f gold\n", item.Nom, item.Quantity, item.Prix)
	}
}
