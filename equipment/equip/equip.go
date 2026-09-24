package equip

import (
	"fmt"
	char "jeu/Character"
	Equipment "jeu/equipment"
)

func Equip(c *char.Character, slot string, newEquip Equipment.Equipment) {
	if c.Equipment == nil {
		c.Equipment = make(map[string]Equipment.Equipment)
	}

	if currentEquip, exists := c.Equipment[slot]; exists {
		c.Health -= currentEquip.Healthbonus
		c.Vitesse -= currentEquip.Speedbonus
		c.Strength -= currentEquip.Damagebonus
		fmt.Printf("Vous avez retiré : %s\n", currentEquip.Name)
	}

	c.Equipment[slot] = newEquip
	c.Health += newEquip.Healthbonus
	c.Vitesse += newEquip.Speedbonus
	c.Strength += newEquip.Damagebonus

	fmt.Printf("Vous avez équipé : %s dans l'emplacement [%s] !\n", newEquip.Name, slot)
}
