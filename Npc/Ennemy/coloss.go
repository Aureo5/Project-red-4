package ennemy

import (
	character "jeu/Character"
	item "jeu/item"
)

var Coloss = character.Character{Health: 150, Strength: 20, Vitesse: 60, XPValue: 30, Or: 25, Inventory: []item.Item{item.Bone(2), item.NewSkin(2)}}
