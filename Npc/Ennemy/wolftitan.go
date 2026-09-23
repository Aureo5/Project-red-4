package ennemy

import (
	character "jeu/Character"
	item "jeu/item"
)

var Wolftitan = character.Character{Health: 120, Strength: 15, Vitesse: 50, Inventory: []item.Item{item.Bone(2), item.NewSkin(3), item.NewWolfSkin(1)}}
