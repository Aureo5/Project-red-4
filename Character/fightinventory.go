package character

import (
	"fmt"
)

func (c *Character) FightInventory() {
	fmt.Printf("\n=== STATUT EN COMBAT ===\n")
	fmt.Printf("Joueur : %s | PV : %d | Force : %d | Vitesse : %d\n", c.Name, c.Health, c.Strength, c.Vitesse)

	var indexConsommables []int

	fmt.Println("\n=== CONSOMMABLES DISPONIBLES ===")
	compteur := 1
	for realIndex, itemObj := range c.Inventory {
		if estConsommable(itemObj.Nom) {
			fmt.Printf("%d. %s (x%d)\n", compteur, itemObj.Nom, itemObj.Quantity)
			indexConsommables = append(indexConsommables, realIndex)
			compteur++
		}
	}

	if len(indexConsommables) == 0 {
		fmt.Println("Vous n'avez aucun consommable utilisable en combat.")
		return
	}

	fmt.Println("0. Annuler (retour au choix d'action)")

	var choix int
	fmt.Print("\nChoisis un consommable à utiliser : ")
	fmt.Scan(&choix)

	if choix <= 0 || choix > len(indexConsommables) {
		return
	}

	targetIndex := indexConsommables[choix-1]
	itemChoisi := c.Inventory[targetIndex]

	switch itemChoisi.Nom {
	case "Potion de soin", "Heal Syringe":
		soin := 50
		c.Health += soin
		fmt.Printf("Vous utilisez %s et récupérez %d PV ! (PV actuels : %d)\n", itemChoisi.Nom, soin, c.Health)
		c.RemoveItemAtIndex(targetIndex)

	default:
		fmt.Println("Impossible d'utiliser cet objet.")
	}
}

func estConsommable(nom string) bool {
	switch nom {
	case "Potion de soin", "Heal Syringe":
		return true
	default:
		return false
	}
}
