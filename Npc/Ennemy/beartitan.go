package ennemy

import (
	character "jeu/Character"
	item "jeu/item"
)

var Beartitan = character.Character{Name: "Geant Ursadon", Health: 220, Strength: 30, Vitesse: 65, XPValue: 50, Or: 45, Inventory: []item.Item{item.Bone(3), item.NewSkin(3), item.NewBearSkin(1)}}
