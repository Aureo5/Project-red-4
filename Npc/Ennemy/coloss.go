package ennemy

import (
	character "jeu/Character"
	item "jeu/item"
)

var coloss = character.Character{Health: 150, Strength: 20, Vitesse: 60, Inventory: []item.Item{item.Bone(2), item.NewSkin(2)}}
