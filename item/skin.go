package item

func NewSkin(qty int) Item {
	return Item{ID: "Skin", Nom: "Peau de monstre", Prix: 10.0, Quantity: qty, Weight: 0.2}
}
