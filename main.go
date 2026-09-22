package main

import (
	character "jeu/Character"
	camp "jeu/camp"
	firstdtrict "jeu/firstdistrict"
)

func main() {
	character.CreateCharacter()
	firstdtrict.Firstdistrict()
	camp.Choicecamp()
}
