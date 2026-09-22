package jeu

func NewHealSyringe(qty int) Item {
	return Item{ID: "heal_syringe", Nom: "Seringue de Soin", Prix: 45.0, Quantity: qty}
}

func HealSyringe() {

}
