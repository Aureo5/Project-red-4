package jeu

import (
	"fmt"
	"jeu/Character"
)

func (c *Character) AddItem(i item.Item) {
	for idx, objet := range c.Inventory {
		if objet.ID == i.ID {
			c.Inventory[idx].Quantity += i.Quantity
			return
		}
	}
	c.Inventory = append(c.Inventory, i)
}

// Acheter vérifie l'or disponible, le déduit, et ajoute l'item à l'inventaire
func (c *Character) buy(i item.Item) {
	coutTotal := i.Prix * float64(i.Quantity)

	if float64(c.Or) < coutTotal {
		fmt.Println("❌ Pas assez d'or pour acheter cet objet.")
		return
	}

	c.Or -= int(coutTotal)
	c.AddItem(i)
	fmt.Printf("✅ %s x%d ajouté à l'inventaire ! (-%.0f or, il te reste %d or)\n", i.Nom, i.Quantity, coutTotal, c.Or)
}
