package character

import (
	"fmt"
	Equipment "jeu/equipment"
)

func (c *Character) EquipFromInventory() {
	var indexValides []int

	fmt.Println("\n=== ÉQUIPEMENTS DISPONIBLES EN INVENTAIRE ===")

	compteur := 1
	for indexReel, itemObj := range c.Inventory {
		if estEquipement(itemObj.Nom) {
			fmt.Printf("%d. %s (x%d)\n", compteur, itemObj.Nom, itemObj.Quantity)
			indexValides = append(indexValides, indexReel)
			compteur++
		}
	}

	if len(indexValides) == 0 {
		fmt.Println("Aucun équipement disponible dans ton inventaire.")
		return
	}

	fmt.Println("0. Retour")

	var choix int
	fmt.Print("\nChoisis un équipement à porter (numéro) : ")
	fmt.Scan(&choix)

	if choix <= 0 || choix > len(indexValides) {
		return
	}

	indexCible := indexValides[choix-1]
	itemChoisi := c.Inventory[indexCible]

	switch itemChoisi.Nom {

	case "Lame basique":
		Equip(c, "lame", Equipment.Basicsword)
		c.RemoveItemAtIndex(indexCible)

	case "Grappling":
		Equip(c, "grappin", Equipment.Grappling)
		c.RemoveItemAtIndex(indexCible)

	case "Propulseur basique":
		Equip(c, "propulseur", Equipment.Basicpropulsor)
		c.RemoveItemAtIndex(indexCible)

	default:
		fmt.Println("Cet objet ne peut pas être équipé.")
	}
}

func estEquipement(nom string) bool {
	switch nom {
	case "Lame basique", "Grappling", "Propulseur basique":
		return true
	default:
		return false
	}
}

func (c *Character) RemoveItemAtIndex(index int) {
	c.Inventory[index].Quantity--
	if c.Inventory[index].Quantity <= 0 {
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	}
}
