package jeu

import ("jeu/Project-red-4/item")

combatsyringe = item{}

func CombatSyringe() {
	
	tour := 3
	Bonus := (strength * 100) % 35
	strength += Bonus
	for tour > 0 {
		tour--
	}
	strength -= Bonus
}
