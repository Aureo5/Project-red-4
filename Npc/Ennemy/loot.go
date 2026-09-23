package ennemy

import (
	"fmt"
	character "jeu/Character"
)

func GiveLoot(player *character.Character, mob character.Character) {
	if len(mob.Inventory) == 0 {
		fmt.Println("Le monstre n'avait rien sur lui...")
		return
	}

	fmt.Println("\n--- BUTIN OBTENU ---")
	for _, lootItem := range mob.Inventory {
		player.AddItem(lootItem)
		fmt.Printf("• %s (x%d)\n", lootItem.Nom, lootItem.Quantity)
	}
}
