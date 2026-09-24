package character

import (
	"fmt"
	item "jeu/item"
)

func (c *Character) AddItem(newItem item.Item) {
	poidsAjoute := newItem.Weight * float64(newItem.Quantity)

	// Vérification de la limite de poids
	if c.GetCurrentWeight()+poidsAjoute > c.MaxWeight {
		fmt.Printf("❌ Trop lourd ! Impossible d'ajouter %s (Poids : %.1f/%.1f kg)\n",
			newItem.Nom, c.GetCurrentWeight()+poidsAjoute, c.MaxWeight)
		return
	}

	// Ajout ou empilement de l'objet
	for i, it := range c.Inventory {
		if it.Nom == newItem.Nom {
			c.Inventory[i].Quantity += newItem.Quantity
			fmt.Printf("Vous avez ajouté %s (Total : %d | Poids : %.1f kg)\n",
				newItem.Nom, c.Inventory[i].Quantity, c.GetCurrentWeight())
			return
		}
	}

	c.Inventory = append(c.Inventory, newItem)
	fmt.Printf("Vous avez obtenu : %s ! (Poids total : %.1f/%.1f kg)\n",
		newItem.Nom, c.GetCurrentWeight(), c.MaxWeight)
}

func (c *Character) Buy(i item.Item) {
	coutTotal := i.Prix * float64(i.Quantity)

	if float64(c.Or) < coutTotal {
		fmt.Println("❌ Pas assez d'or pour acheter cet objet.")
		return
	}

	c.Or -= int(coutTotal)
	c.AddItem(i)
	fmt.Printf("✅ %s x%d ajouté à l'inventaire ! (-%.0f or, il te reste %d or)\n", i.Nom, i.Quantity, coutTotal, c.Or)
}

func (c *Character) GetCurrentWeight() float64 {
	var total float64 = 0
	for _, it := range c.Inventory {
		total += it.Weight * float64(it.Quantity)
	}
	return total
}
