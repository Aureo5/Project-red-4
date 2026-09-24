package ennemy

import (
	character "jeu/Character"
	item "jeu/item"
)

var Giant = character.Character{Name: "Géant", Health: 75, Strength: 10, Vitesse: 45, XPValue: 20, Or: 15, Inventory: []item.Item{item.Bone(1), item.NewSkin(1)}}
