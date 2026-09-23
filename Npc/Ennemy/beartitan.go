package ennemy

import (
	character "jeu/Character"
	item "jeu/item"
)

var beartitan = character.Character{Health: 220, Strength: 30, Vitesse: 65, Inventory: []item.Item{item.Bone(3), item.NewSkin(3), item.NewBearSkin(1)}}
