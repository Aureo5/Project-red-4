package jeu

func NewBearSkin(qty int) Item {
	return Item{ID: "bear_skin", Nom: "peau d'ours", Prix: 150.0, Quantity: qty}
}
