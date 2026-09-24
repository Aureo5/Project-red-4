package removeitem

import (
	char "jeu/Character"
)

func RemoveItemAtIndex(c *char.Character, index int) {
	c.Inventory[index].Quantity--
	if c.Inventory[index].Quantity <= 0 {
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	}
}
