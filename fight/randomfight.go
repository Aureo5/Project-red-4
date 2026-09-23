package fight

import (
	character "jeu/Character"
	"jeu/Npc/Ennemy"
	"math/rand"
)

func LaunchRandomFight(perso *character.Character) {
	// Tirage d'un nombre : 0 ou 1
	randomChoice := rand.Intn(2)

	var currentMob *character.Character

	if randomChoice == 0 {
		currentMob = &ennemy.Giant
	} else {
		currentMob = &ennemy.Beartitan
	}

	// Lancement du combat avec l'ennemi sélectionné
	StartFight(perso, currentMob)
}
