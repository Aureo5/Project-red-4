package fight

import (
	character "jeu/Character"
	"jeu/Npc/Ennemy"
	"math/rand"
)

func LaunchRandomFight(perso *character.Character) {
	randomChoice := rand.Intn(2)

	var currentMob *character.Character

	if randomChoice == 0 {
		currentMob = &ennemy.Giant
	} else {
		currentMob = &ennemy.Beartitan
	}

	StartFight(perso, currentMob)
}
