package character

import (
	"fmt"
	equip "jeu/equipment"
)

func (c *Character) AccessInventory() {
	if len(c.Inventory) == 0 {
		fmt.Println("\nVotre inventaire est vide.")
		return
	}

	for {
		if len(c.Inventory) == 0 {
			fmt.Println("\nVotre inventaire est désormais vide.")
			return
		}

		fmt.Println("\n=== INVENTAIRE ===")
		for i, it := range c.Inventory {
			fmt.Printf("%d. %s (x%d)\n", i+1, it.Nom, it.Quantity)
		}
		fmt.Println("0. Retour")

		var choix int
		fmt.Print("\nChoisis un objet à utiliser/équiper (numéro) : ")
		fmt.Scan(&choix)

		if choix == 0 {
			return
		}

		if choix < 1 || choix > len(c.Inventory) {
			fmt.Println("Choix invalide.")
			continue
		}

		itemChoisi := c.Inventory[choix-1]

		// Gestion selon le type/propriétés de l'objet
		switch itemChoisi.Type {

		case "potion", "soin", "consommable":
			soin := 50
			c.Health += soin
			fmt.Printf("Vous avez consommé %s (+%d PV) ! PV actuels : %d\n", itemChoisi.Nom, soin, c.Health)

			// Retrait / Décrémentation
			c.RemoveItemAtIndex(choix - 1)

		case "equipement", "lame", "grappin", "propulseur":
			slot := itemChoisi.Slot
			if slot == "" {
				slot = "lame"
			}
			nouveauMatos := equip.Equipment{
				Name:        itemChoisi.Nom,
				Healthbonus: itemChoisi.Healthbonus,
				Speedbonus:  itemChoisi.Speedbonus,
				Damagebonus: itemChoisi.Damagebonus,
			}

			Equip(c, slot, nouveauMatos)

			c.RemoveItemAtIndex(choix - 1)

		default:
			fmt.Printf("Impossible d'utiliser l'objet %s directement.\n", itemChoisi.Nom)
		}
	}
}

func (c *Character) RemoveItemAtIndex(index int) {
	c.Inventory[index].Quantity--
	if c.Inventory[index].Quantity <= 0 {
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	}
}
