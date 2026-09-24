package ennemy

import (
	character "jeu/Character"
	item "jeu/item"
)

var Wolftitan = character.Character{Name: "Titan gueule de loup", Health: 120, Strength: 15, Vitesse: 50, XPValue: 30, Or: 35, Inventory: []item.Item{item.Bone(2), item.NewSkin(3), item.NewWolfSkin(1)}, BurnTurns: 0}
