package interactiveInventory

import (
	"fmt"
	char "jeu/Character"
	add "jeu/Character/Additem"
	display "jeu/Character/inventory/inventorydisplay"
)

// À utiliser hors combat uniquement
func DisplayCharacterMenu(c *char.Character) {
	for {
		fmt.Println("\n=== GESTION DU PERSONNAGE ===")
		fmt.Println("1. Voir les informations du personnage")
		fmt.Println("2. Voir les équipements équipés")
		fmt.Println("3. Voir tout l'inventaire")
		fmt.Println("0. Retour au jeu")

		var choix string
		fmt.Print("Choix : ")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			display.Displayinfo(*c)
		case "2":
			DisplayEquipment(c)
		case "3":
			DisplayFullInventory(c)
		case "0":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func DisplayEquipment(c *char.Character) {
	fmt.Println("\n=== ÉQUIPEMENTS ===")
	if len(c.Equipment) == 0 {
		fmt.Println("Aucun équipement actuellement porté.")
		return
	}
	for slot, eq := range c.Equipment {
		fmt.Printf("• [%s] : %s (PV: +%d | Dégâts: +%d | Vitesse: +%d)\n",
			slot, eq.Name, eq.Healthbonus, eq.Damagebonus, eq.Speedbonus)
	}
}

func DisplayFullInventory(c *char.Character) {
	fmt.Println("\n=== INVENTAIRE ===")
	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}
	fmt.Printf("%.1fkg / %.1fkg\n", add.GetCurrentWeight(c), c.MaxWeight)
	for i, it := range c.Inventory {
		fmt.Printf("%d. %s (x%d)\n", i+1, it.Nom, it.Quantity)
	}
}
