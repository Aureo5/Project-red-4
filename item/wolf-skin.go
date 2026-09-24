package item

func NewWolfSkin(qty int) Item {
	return Item{ID: "wolf_skin", Nom: "Peau de loup", Prix: 150.0, Quantity: qty, Weight: 0.8}
}
