package jeu

func NewBearSkin(qty int) Item {
	return Item{ID: "bear_skin", Nom: "peau d'ours", Prix: 300.0, Quantity: qty}
}
