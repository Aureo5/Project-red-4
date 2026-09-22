package jeu

import (
	"fmt"
	item "jeu/item"
)

type Character struct {
	Name       string      // nom du personnage
	Health     int         // points de vie du personnage
	Vitesse    int         // vitesse du personnage
	Experience int         // points d'expérience du personnage
	Nameclass  string      // classe du personnage
	Strength   int         // force du personnage
	Or         int         // or du personnage
	Inventory  []item.Item //inventaire du personnage
}

// vérifie si le nom est valide (première lettre majusculeS, reste en minuscule)
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

// redemande le nom jusqu'à que l'utilisateur entre un nom valide
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
func (c Character) Displayinfo() {
	fmt.Println("\n===Informations sur le personnage=== ")
	fmt.Printf("Nom        : %s\n", c.Name)
	fmt.Printf("Classe     : %s\n", c.Nameclass)
	fmt.Printf("Vie        : %d\n", c.Health)
	fmt.Printf("Force      : %d\n", c.Strength)
	fmt.Printf("Vitesse    : %d\n", c.Vitesse)
	fmt.Printf("Expérience : %d\n", c.Experience)
	fmt.Printf("Or         : %d\n", c.Or)
	fmt.Println("======================================")
}
