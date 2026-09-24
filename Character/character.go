package character

import (
	"fmt"
	equip "jeu/equipment"
	item "jeu/item"
)

type Character struct {
	Name           string
	Health         int
	Vitesse        int
	Experience     int
	Level          int
	XPValue        int
	Nameclass      string
	Strength       int
	Or             int
	Inventory      []item.Item
	MaxWeight      float64
	Equipment      map[string]equip.Equipment
	SkillCooldowns map[string]int

	BurnTurns   int
	DamageBoost int

	TurnCounter int
	IsDefending bool
}

func Maj(nom string) bool {
	if len(nom) == 0 {
		return false
	}

	for i, lettre := range nom {
		if i == 0 {
			if lettre < 'A' || lettre > 'Z' {
				return false
			}
		} else {
			if lettre < 'a' || lettre > 'z' {
				return false
			}
		}
	}

	return true
}

func Name() string {
	var nom string

	for {
		fmt.Print("\n===Choisis le nom de ton personnage=== \n⚠️  le nom doit contenir une majuscule au début puis des minuscules et aucun caractère spécial: ")
		fmt.Scan(&nom)

		if Maj(nom) {
			return nom
		}

		fmt.Println("Nom invalide, réessaie.")
	}
}

func (c *Character) HasResource(id string, qty int) bool {
	for _, it := range c.Inventory {
		if it.ID == id {
			return it.Quantity >= qty
		}
	}
	return false
}

func (c *Character) ConsumeResource(id string, qty int) {
	for i, it := range c.Inventory {
		if it.ID == id {
			c.Inventory[i].Quantity -= qty
			return
		}
	}
}

func (c *Character) Upgrade(key string, peauCost, osCost, orCost int, apply func(*equip.Equipment)) bool {
	if !c.HasResource("Skin", peauCost) {
		fmt.Println("Pas assez de peau pour cette amélioration.")
		return false
	}
	if !c.HasResource("bone", osCost) {
		fmt.Println("Pas assez d'os pour cette amélioration.")
		return false
	}
	if c.Or < orCost {
		fmt.Println("Pas assez d'or pour cette amélioration.")
		return false
	}

	c.ConsumeResource("Skin", peauCost)
	c.ConsumeResource("bone", osCost)
	c.Or -= orCost

	eq := c.Equipment[key]
	apply(&eq)
	c.Equipment[key] = eq

	fmt.Printf("%s amélioré avec succès !\n", eq.Name)
	return true
}

func (c *Character) UpgradeSpecial(key string, specialID string, specialNom string, specialQty int, peauCost, osCost, orCost int, apply func(*equip.Equipment)) bool {
	if !c.HasResource(specialID, specialQty) {
		fmt.Printf("Pas assez de %s pour cette amélioration.\n", specialNom)
		return false
	}
	if !c.HasResource("Skin", peauCost) {
		fmt.Println("Pas assez de peau pour cette amélioration.")
		return false
	}
	if !c.HasResource("bone", osCost) {
		fmt.Println("Pas assez d'os pour cette amélioration.")
		return false
	}
	if c.Or < orCost {
		fmt.Println("Pas assez d'or pour cette amélioration.")
		return false
	}

	c.ConsumeResource(specialID, specialQty)
	c.ConsumeResource("Skin", peauCost)
	c.ConsumeResource("bone", osCost)
	c.Or -= orCost

	eq := c.Equipment[key]
	apply(&eq)
	c.Equipment[key] = eq

	fmt.Printf("%s amélioré avec succès !\n", eq.Name)
	return true
}

func (c *Character) GiveXP(amount int) {
	if c.Level == 0 {
		c.Level = 1
	}

	c.Experience += amount
	fmt.Printf("✨ Vous gagnez %d XP ! (%d/%d)\n", amount, c.Experience, c.Level*100)

	for c.Experience >= c.Level*100 {
		c.Experience -= c.Level * 100
		c.LevelUp()
	}
}

func (c *Character) LevelUp() {
	c.Level++
	c.Health += 10
	c.Strength += 3
	fmt.Printf("===NIVEAU SUPÉRIEUR=== \nVous êtes maintenant niveau %d ! (+20 PV max, +5 Force)\n", c.Level)
}
