package ennemy

import (
	"fmt"
	character "jeu/Character"
)

func EnemyTurn(mob *character.Character, target *character.Character) {
	mob.TurnCounter++
	mob.IsDefending = false

	if mob.TurnCounter%2 != 0 {
		fmt.Printf("⚔️ %s attaque %s et inflige %d dégâts !\n", mob.Name, target.Name, mob.Strength)
		target.Health -= mob.Strength
	} else {
		mob.IsDefending = true
		fmt.Printf("🛡️ %s se met en posture défensive ! (Dégâts subis réduits de 50%% au prochain tour)\n", mob.Name)
	}
}
