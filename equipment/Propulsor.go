package jeu

func NewPropulsor() Equipment {
	return Equipment{ID: "basic_propulsor", Name: "Propulseur", Gaz: 100, Healthbonus: 10, Speedbonus: 5}
}
