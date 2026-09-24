package character

import (
	"fmt"
)

// À utiliser hors combat uniquement
func (c *Character) DisplayCharacterMenu() {
	for {
		fmt.Println("\n=== GESTION DU PERSONNAGE ===")
		fmt.Println("1. Voir les informations du personnage")
		fmt.Println("2. Voir les équipements équipés")
		fmt.Println("3. Changer d'équipement (depuis l'inventaire)")
		fmt.Println("4. Voir tout l'inventaire")
		fmt.Println("0. Retour au jeu")

		var choix string
		fmt.Print("Choix : ")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			Displayinfo(*c)
		case "2":
			c.DisplayEquipment()
		case "3":
			c.EquipFromInventory()
		case "4":
			c.DisplayFullInventory()
		case "0":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (c *Character) DisplayEquipment() {
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

func (c *Character) DisplayFullInventory() {
	fmt.Println("\n=== INVENTAIRE ===")
	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}
	fmt.Printf("%.1fkg / %.1fkg\n", c.GetCurrentWeight(), c.MaxWeight)
	for i, it := range c.Inventory {
		fmt.Printf("%d. %s (x%d)\n", i+1, it.Nom, it.Quantity)
	}
}
