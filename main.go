package main

import (
	character "jeu/Character"
	Menu "jeu/Menu"
	camp "jeu/camp"
	firstdtrict "jeu/firstdistrict"
)

func main() {
	Menu.Menu()
	character.CreateCharacter()
	firstdtrict.Firstdistrict()
	camp.Choicecamp()
}
