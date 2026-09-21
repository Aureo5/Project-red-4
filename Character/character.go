package jeu

import (
	"fmt"
)

type Character struct {
	Name       string // nom du personnage
	health     int    // points de vie du personnage
	vitesse    int    // vitesse du personnage
	experience int    // points d'expérience du personnage
	nameclass  string // classe du personnage
	strength   int    // force du personnage
	or         int    // or du personnage
}

// vérifie si le nom est valide (première lettre majuscule, reste en minuscule)
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
		fmt.Print("===Choisis le nom de ton personnage=== \n⚠️  le nom doit contenir une majuscule au début puis des minuscules: ")
		fmt.Scan(&nom)

		if Maj(nom) {
			return nom
		}

		fmt.Println("Nom invalide, réessaie.")
	}
}
func (c Character) Info() {
	fmt.Println("===Informations sur le personnage=== :")
	fmt.Printf("Nom        : %s\n", c.Name)
	fmt.Printf("Classe     : %s\n", c.nameclass)
	fmt.Printf("Vie        : %d\n", c.health)
	fmt.Printf("Force      : %d\n", c.strength)
	fmt.Printf("Vitesse    : %d\n", c.vitesse)
	fmt.Printf("Expérience : %d\n", c.experience)
	fmt.Printf("Or         : %d\n", c.or)
	fmt.Println("======================================")
}
