package character

import (
	"fmt"
	equip "jeu/equipment"
	item "jeu/item"
)

type Character struct {
	Name       string
	Health     int
	Vitesse    int
	Experience int
	Nameclass  string
	Strength   int
	Or         int
	Inventory  []item.Item
	Equipment  map[string]equip.Equipment
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
