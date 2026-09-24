package ennemy

import (
	"fmt"
	character "jeu/Character"
)

func GiveLoot(player *character.Character, mob character.Character) {
	if mob.XPValue > 0 {
		player.GiveXP(mob.XPValue)
	}
	if len(mob.Inventory) == 0 {
		fmt.Println("Le monstre n'avait rien sur lui...")
		return
	}

	if mob.Or > 0 {
		player.Or += mob.Or
		fmt.Printf(" Vous gagnez %d or ! (total : %d)\n", mob.Or, player.Or)
	}
	fmt.Println("\n===BUTIN OBTENU===")
	for _, lootItem := range mob.Inventory {
		player.AddItem(lootItem)
		fmt.Printf("• %s (x%d)\n", lootItem.Nom, lootItem.Quantity)
	}
}
