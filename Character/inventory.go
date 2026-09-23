package character

import (
	"fmt"
)

func DisplayInv(c Character) {
	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}
	fmt.Println("--- INVENTAIRE ---")
	for _, item := range c.Inventory {
		fmt.Printf("- %s (x%d) : %.2f gold\n", item.Nom, item.Quantity, item.Prix)
	}
}

func DisplayEquip(c Character) {
	if len(c.Equipment) == 0 {
		fmt.Println("Votre inventaire d'équipement est vide.")
		return
	}
	fmt.Println("--- Equipement ---")
	for _, item := range c.Equipment {
		fmt.Printf("Lame: %s \n", item.Name)
	}
}
