package jeu

type burn struct {
	tour   int
	damage int
}

func NewGrenadeIncendiary(qty int) Item {
	return Item{ID: "Incendiary_Grenade", Nom: "Grenade incendiaire", Prix: 40.0, Quantity: qty}
}
