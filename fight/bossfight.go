package fight

import (
	"fmt"
	character "jeu/Character"
	"jeu/Npc/Ennemy"
	"jeu/skill"
)

func BossFight(perso *character.Character, mob *character.Character) {
	fmt.Printf("\nUN COMBAT COMMENCE CONTRE : %s (PV: %d) ⚔️\n", mob.Name, mob.Health)
	save := mob.Health
	saveperso := perso.Health
	tour := 1

	for perso.Health > 0 && mob.Health > 0 {
		fmt.Printf("\n--- TOUR %d ---\n", tour)
		if mob.BurnTurns > 0 {
			mob.Health -= 10
			mob.BurnTurns--
			fmt.Printf("Le %s subit 10 dégâts de brûlure ! (PV restants : %d)\n", mob.Name, mob.Health)
		}

		fmt.Printf("Vos PV: %d | PV de %s: %d\n", perso.Health, mob.Name, mob.Health)
		fmt.Println("1. Attaquer")
		fmt.Println("2. Compétence spéciale")
		fmt.Println("3. Ouvrir l'inventaire")
		fmt.Println("4. Fuir")

		var choix string
		fmt.Print("Action : ")
		fmt.Scan(&choix)
		actionUtilisee := true

		switch choix {
		case "1":
			degatsJoueur := perso.Strength
			if perso.DamageBoost > 0 {
				degatsJoueur *= 2
				perso.DamageBoost--
				fmt.Println("Attaque boostée de +100% !")
			}
			mob.Health -= degatsJoueur
			fmt.Printf("Vous infligez %d dégâts à %s !\n", degatsJoueur, mob.Name)

		case "2":
			actionUtilisee = skill.UseSkill(perso, mob)
		case "3":
			perso.FightInventory(saveperso)

		case "4":
			fmt.Println("Vous avez pris la fuite !")
			return

		default:
			fmt.Println("Choix invalide, vous passez votre tour !")
		}

		if mob.Health <= 0 {
			fmt.Printf("\nVous avez vaincu %s !\n", mob.Name)
			ennemy.GiveLoot(perso, *mob)
			perso.Health = saveperso
			mob.Health = save
			perso.Unlock = true
			return
		}

		if actionUtilisee {
			skill.TickCooldowns(perso)
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
