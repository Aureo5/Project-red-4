package fight

import (
	character "jeu/Character"
	"jeu/Npc/Ennemy"
	fight "jeu/fight"
	"math/rand"
)

func LaunchRandomFight(perso *character.Character) {
	randomChoice := rand.Intn(2)

	var currentMob *character.Character

	if randomChoice == 0 {
		currentMob = &ennemy.Giant
	} else {
		currentMob = &ennemy.Coloss
	}

	fight.StartFight(perso, currentMob)
}
