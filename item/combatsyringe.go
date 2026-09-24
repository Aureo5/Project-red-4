package item

func NewCombatSyringe(qty int) Item {
	return Item{ID: "combat_syringe", Nom: "Seringue de Combat", Prix: 55.0, Quantity: qty, Weight: 0.4, Type: "Consumable"}
}
