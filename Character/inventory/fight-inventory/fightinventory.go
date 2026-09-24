package fightinventory

import (
	"fmt"
	char "jeu/Character"
)

func FightInventory(c *char.Character, maxhealth int) {
	fmt.Printf("\n=== STATUT ===\n")
	fmt.Printf("Joueur : %s | PV : %d | Force : %d | Vitesse : %d\n", c.Name, c.Health, c.Strength, c.Vitesse)

	var indexConsommables []int

	fmt.Println("\n=== CONSOMMABLES DISPONIBLES ===")
	compteur := 1
	for realIndex, itemObj := range c.Inventory {
		if itemObj.Type == "consommable" {
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
	fmt.Print("\nChoisis un consommable à utiliser (lettre par lettre) : ")
	fmt.Scan(&choix)

	if choix <= 0 || choix > len(indexConsommables) {
		return
	}

	targetIndex := indexConsommables[choix-1]
	itemChoisi := c.Inventory[targetIndex]

	switch itemChoisi.ID {

	case "grenade_incendiaire":
		fmt.Printf("Vous lancez une %s ! L'ennemi prendra 10 dégâts de brûlure par tour.\n", itemChoisi.Nom)
		c.RemoveItemAtIndex(targetIndex)

	case "seringue_combat":
		c.DamageBoost = 2
		fmt.Printf("Vous utilisez une %s ! Vos dégâts sont augmentés de 100%% pendant 2 tours.\n", itemChoisi.Nom)
		c.RemoveItemAtIndex(targetIndex)

	case "seringue_soin":
		heal := 50
		if c.Health > (maxhealth - heal) {
			c.Health += (maxhealth - c.Health)
			heal = (maxhealth - c.Health)
		} else {
			c.Health += heal
		}
		fmt.Printf("Vous utilisez %s et récupérez %d PV ! (PV actuels : %d)\n", itemChoisi.Nom, heal, c.Health)
		c.RemoveItemAtIndex(targetIndex)
	default:
		return
	}
}
