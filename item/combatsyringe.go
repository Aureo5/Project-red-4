package item

func NewCombatSyringe(qty int) Item {
	return Item{ID: "combat_syringe", Nom: "Seringue de Combat", Prix: 25.0, Quantity: qty}
}
