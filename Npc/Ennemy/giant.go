package ennemy

import (
	character "jeu/Character"
	item "jeu/item"
)

var giant = character.Character{Health: 75, Strength: 10, Vitesse: 45, Inventory: []item.Item{item.Bone(1), item.NewSkin(1)}}
