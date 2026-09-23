package item

func NewHealSyringe(qty int) Item {
	return Item{ID: "heal_syringe", Nom: "Seringue de Soin", Prix: 15.0, Quantity: qty}
}
