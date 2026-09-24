package character

import (
	"fmt"
)

func Displayinfo(c Character) {
	fmt.Println("\n===Informations sur le personnage=== ")
	fmt.Printf("Nom        : %s\n", c.Name)
	fmt.Printf("Classe     : %s\n", c.Nameclass)
	fmt.Printf("Vie        : %d\n", c.Health)
	fmt.Printf("Force      : %d\n", c.Strength)
	fmt.Printf("Vitesse    : %d\n", c.Vitesse)
	fmt.Printf("Expérience : %d\n", c.Experience)
	fmt.Printf("Or         : %d\n", c.Or)
	fmt.Println("======================================")
}

func DisplayInv(c Character) {
	fmt.Println("\n=== INVENTAIRE ===")
	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}
	for _, item := range c.Inventory {
		fmt.Printf("- %s (x%d) : %.2f or\n", item.Nom, item.Quantity, item.Prix)
	}
}

func DisplayEquip(c Character) {
	fmt.Println("\n=== Equipement ===")
	if len(c.Equipment) == 0 {
		fmt.Println("Votre inventaire d'équipement est vide.")
		return
	}
	for _, item := range c.Equipment {
		fmt.Printf("%s \n", item.Name)
	}
}
