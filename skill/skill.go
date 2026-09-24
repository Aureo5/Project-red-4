package skill

import (
	"fmt"
	char "jeu/Character"
)

var bossNames = []string{"Geant Ursadon", "Titan gueule de loup", "Titan colossale"}

func isBoss(mob *char.Character) bool {
	for _, nom := range bossNames {
		if mob.Name == nom {
			return true
		}
	}
	return false
}
func SkillReady(c *char.Character, name string) bool {
	if c.SkillCooldowns == nil {
		return true
	}
	return c.SkillCooldowns[name] <= 0
}

func TriggerCooldown(c *char.Character, name string) {
	if c.SkillCooldowns == nil {
		c.SkillCooldowns = make(map[string]int)
	}
	c.SkillCooldowns[name] = 3
}

func TickCooldowns(c *char.Character) {
	for name, turns := range c.SkillCooldowns {
		if turns > 0 {
			c.SkillCooldowns[name] = turns - 1
		}
	}
}

func UseSkill(c *char.Character, mob *char.Character) bool {
	switch c.Nameclass {
	case "eclaireur":
		return useSkillEclaireur(c, mob)
	case "soldat":
		return useSkillSoldat(c, mob)
	case "medic":
		return useSkillMedic(c, mob)
	default:
		fmt.Println("Ta classe n'a aucune compétence spéciale.")
		return false
	}
}

func useSkillEclaireur(c *char.Character, mob *char.Character) bool {
	fmt.Println("\n=== COMPÉTENCES ÉCLAIREUR ===")
	if SkillReady(c, "coupe_membres") {
		fmt.Println("1. Coupe les membres (+5 dégâts) - Disponible")
	} else {
		fmt.Printf("1. Coupe les membres - En recharge (%d tour(s))\n", c.SkillCooldowns["coupe_membres"])
	}
	if SkillReady(c, "fusee") {
		fmt.Println("2. Fusée (+20 dégâts) - Disponible")
	} else {
		fmt.Printf("2. Fusée - En recharge (%d tour(s))\n", c.SkillCooldowns["fusee"])
	}
	fmt.Println("0. Annuler")

	var choix string
	fmt.Print("Choix : ")
	fmt.Scan(&choix)

	switch choix {
	case "1":
		if !SkillReady(c, "coupe_membres") {
			fmt.Println("Compétence en recharge.")
			return false
		}
		degats := c.Strength + 5
		mob.Health -= degats
		fmt.Printf("Vous coupez les membres de %s et infligez %d dégâts !\n", mob.Name, degats)
		TriggerCooldown(c, "coupe_membres")
		return true
	case "2":
		if !SkillReady(c, "fusee") {
			fmt.Println("Compétence en recharge.")
			return false
		}
		degats := c.Strength + 20
		mob.Health -= degats
		fmt.Printf("Vous appelez un soldat en renfort et infligez %d dégâts !\n", degats)
		TriggerCooldown(c, "fusee")
		return true
	default:
		fmt.Println("Annulé.")
		return false
	}
}

func useSkillSoldat(c *char.Character, mob *char.Character) bool {
	fmt.Println("\n=== COMPÉTENCES SOLDAT ===")
	if SkillReady(c, "recharge") {
		fmt.Println("1. Changement de lame & gaz de recharge (restaure durabilité/gaz équipés) - Disponible")
	} else {
		fmt.Printf("1. Changement de lame & gaz de recharge - En recharge (%d tour(s))\n", c.SkillCooldowns["recharge"])
	}
	if SkillReady(c, "dechainement") {
		fmt.Println("2. Déchaînement (-30% gaz/durabilité, +30 dégâts) - Disponible")
	} else {
		fmt.Printf("2. Déchaînement - En recharge (%d tour(s))\n", c.SkillCooldowns["dechainement"])
	}
	fmt.Println("0. Annuler")

	var choix string
	fmt.Print("Choix : ")
	fmt.Scan(&choix)

	switch choix {
	case "1":
		if !SkillReady(c, "recharge") {
			fmt.Println("Compétence en recharge.")
			return false
		}
		for slot, eq := range c.Equipment {
			eq.Durability += eq.Durability * 50 / 100
			eq.Gaz += eq.Gaz * 50 / 100
			c.Equipment[slot] = eq
		}
		fmt.Println("Vous rechargez votre équipement, durabilité et gaz restaurés !")
		TriggerCooldown(c, "recharge")
		return true
	case "2":
		if !SkillReady(c, "dechainement") {
			fmt.Println("Compétence en recharge.")
			return false
		}
		for slot, eq := range c.Equipment {
			eq.Durability -= eq.Durability * 30 / 100
			eq.Gaz -= eq.Gaz * 30 / 100
			c.Equipment[slot] = eq
		}
		degats := c.Strength + 30
		mob.Health -= degats
		fmt.Printf("Vous vous déchaînez et infligez %d dégâts à %s ! (-30%% gaz/durabilité de votre équipement)\n", degats, mob.Name)
		TriggerCooldown(c, "dechainement")
		return true
	default:
		fmt.Println("Annulé.")
		return false
	}
}

func useSkillMedic(c *char.Character, mob *char.Character) bool {
	if isBoss(mob) {
		fmt.Println("\n=== COMPÉTENCES MÉDECIN ===")
		if SkillReady(c, "soin") {
			fmt.Println("1. Se soigner (+20 PV) - Disponible")
		} else {
			fmt.Printf("1. Se soigner - En recharge (%d tour(s))\n", c.SkillCooldowns["soin"])
		}
		fmt.Println("2. Appel aux troupes")
		fmt.Println("0. Annuler")

		var choix string
		fmt.Print("Choix : ")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			if !SkillReady(c, "soin") {
				fmt.Println("Compétence en recharge.")
				return false
			}
			c.Health += 20
			fmt.Printf("Vous vous soignez et récupérez 20 PV ! (PV actuels : %d)\n", c.Health)
			TriggerCooldown(c, "soin")
			return true
		case "2":
			// EASTER EGG : kill instantané, pas de TriggerCooldown -> réutilisable sans limite
			mob.Health = 0
			fmt.Printf(" Vous appelez les troupes d'Erwin ! Ils s'occupent du %s, le boss est vaincu !\n", mob.Name)
			return true
		default:
			fmt.Println("Annulé.")
			return false
		}
	}

	// Cas normal (ennemi non-boss) : uniquement le soin, comportement inchangé
	if !SkillReady(c, "soin") {
		fmt.Printf("Compétence en recharge (%d tour(s) restant(s)).\n", c.SkillCooldowns["soin"])
		return false
	}
	c.Health += 20
	fmt.Printf("Vous vous soignez et récupérez 20 PV ! (PV actuels : %d)\n", c.Health)
	TriggerCooldown(c, "soin")
	return true
}
