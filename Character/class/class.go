package character

import (
	"fmt"
	char "jeu/Character"
	Firstdistrict "jeu/district/firstdistrict"
	Equipment "jeu/equipment"
)

// différebntes classes de personnages avec leurs caractéristiques
func Class(nom string) char.Character {
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

		// création du personnage en fonction du choix de l'utilisateur
		switch choice {
		case "1":
			perso := char.Character{Name: nom, Health: 80, Vitesse: 75, Nameclass: "eclaireur", Strength: 15, Or: 100, MaxWeight: 5}
			char.Equip(&perso, "lame", Equipment.Basicsword)
			char.Equip(&perso, "grappin", Equipment.Grappling)
			char.Equip(&perso, "propulseur", Equipment.Basicpropulsor)
			char.Displayinfo(perso)
			Firstdistrict.StartingPoint(&perso)
		case "2":
			perso := char.Character{Name: nom, Health: 100, Vitesse: 50, Nameclass: "soldat", Strength: 20, Or: 100, MaxWeight: 5}
			char.Equip(&perso, "lame", Equipment.Basicsword)
			char.Equip(&perso, "grappin", Equipment.Grappling)
			char.Equip(&perso, "propulseur", Equipment.Basicpropulsor)
			char.Displayinfo(perso)
			Firstdistrict.StartingPoint(&perso)
		case "3":
			perso := char.Character{Name: nom, Health: 120, Vitesse: 50, Nameclass: "medic", Strength: 15, Or: 100, MaxWeight: 5}
			char.Equip(&perso, "lame", Equipment.Basicsword)
			char.Equip(&perso, "grappin", Equipment.Grappling)
			char.Equip(&perso, "propulseur", Equipment.Basicpropulsor)
			char.Displayinfo(perso)
			Firstdistrict.StartingPoint(&perso)
		default:
			fmt.Println("Choix invalide, réessaie.")
		}
	}
}
func CreateCharacter() char.Character {
	// création du personnage
	nom := char.Name()
	perso := Class(nom)
	char.Displayinfo(perso)
	return perso
}
