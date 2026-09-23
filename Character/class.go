package jeu

import (
	"fmt"
	Equipement "jeu/equipment"
	Firstdistrict "jeu/firstdistrict"
)

func Class(nom string) Character {
	var choice string
	for {
		fmt.Println("\n===Choisis ta classe=== ")
		fmt.Println("\n1. Éclaireur - Vie : 80 / Force : 15 / Vitesse : 75% ")
		fmt.Println("compétences spéciales : coupe les membres pour déstabiliser l'ennemi -> +5 points de dégats / fusée -> appelle un soldat +20 points de \ndégats")
		fmt.Println("\n2. Soldat - Vie  : 100 / Force : 20 / Vitesse : 50%")
		fmt.Println("compétences spéciales : changement de lame & gaz de recharge -> redonne instantanément de la durabilité à votre équipement / déchaînement \n-> -30% gaz et durabilité & +30 points de dégats")
		fmt.Println("\n3. Médecin - Vie  : 120 / Force : 15 / Vitesse : 50%")
		fmt.Println("compétences spéciales : possibilités de ce soigner -> +20 points de vies")
		fmt.Println("\n⚠️  Les compétences spéciales sont disponibles tout les 3 tours.")
		fmt.Print("Ton choix (1/2/3) : ")
		fmt.Scan(&choice)

		switch choice {
		case "1":
			perso := Character{Name: nom, Health: 80, Vitesse: 75, Nameclass: "eclaireur", Strength: 15, Or: 100}
			perso.Displayinfo()
			Firstdistrict.StartingPoint()
		case "2":
			perso := Character{Name: nom, Health: 100, Vitesse: 50, Nameclass: "soldat", Strength: 20, Or: 100}
			perso.Displayinfo()
			Firstdistrict.StartingPoint()
		case "3":
			perso := Character{Name: nom, Health: 120, Vitesse: 50, Nameclass: "medic", Strength: 15, Or: 100}
			perso.Displayinfo()
			Firstdistrict.StartingPoint()
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
func CreateCharacter() Character {
	// création du personnage
	nom := Name()
	perso := Class(nom)
	perso.Displayinfo()
	return perso
}
