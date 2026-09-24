package fight

import (
	"fmt"
	character "jeu/Character"
)

func TrainingFight(perso *character.Character, mob *character.Character) {
	fmt.Printf("\nUN COMBAT COMMENCE CONTRE : %s (PV: %d) ⚔️\n", mob.Name, mob.Health)
	save := mob.Health
	saveperso := perso.Health
	tour := 1

	for perso.Health > 0 && mob.Health > 0 {
		fmt.Printf("\n--- TOUR %d ---\n", tour)
		fmt.Printf("Vos PV: %d | PV de %s: %d\n", perso.Health, mob.Name, mob.Health)
		fmt.Println("1. Attaquer")
		fmt.Println("2. Ouvrir l'inventaire")
		fmt.Println("3. Fuir")

		var choix string
		fmt.Print("Action : ")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			degatsJoueur := perso.Strength
			mob.Health -= degatsJoueur
			fmt.Printf("Vous infligez %d dégâts à %s !\n", degatsJoueur, mob.Name)

		case "2":
			perso.FightInventory(saveperso)

		case "3":
			fmt.Println("Vous avez pris la fuite !")
			return

		default:
			fmt.Println("Choix invalide, vous passez votre tour !")
		}

		if mob.Health <= 0 {
			fmt.Printf("\nVous avez vaincu %s !\n", mob.Name)
			perso.Health = saveperso
			mob.Health = save
			return
		}

		fmt.Printf("\nTour de %s...\n", mob.Name)
		perso.Health -= mob.Strength
		fmt.Printf(" %s vous inflige %d dégâts !\n", mob.Name, mob.Strength)

		if perso.Health <= 0 {
			fmt.Println("\nVous avez été vaincu...")
			perso.Health = saveperso
			mob.Health = save
			return
		}

		tour++
	}
}
