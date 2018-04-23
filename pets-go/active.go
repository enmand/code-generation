package main

import (
	"fmt"
	generated "github.com/enmand/code-generation/pets-go/generated"
)

func main() {
	var speaker string
	dog := generated.NewDog()
	speaker = dog.Speak()
	fmt.Printf("a Dog says: %s\n", speaker)
	if err := dog.Walk(); err != nil {
		fmt.Println("Dog are walkable!")
	} else {
		fmt.Println("Dog are not walkable!")
	}
	cat := generated.NewCat()
	speaker = cat.Speak()
	fmt.Printf("a Cat says: %s\n", speaker)
	if err := cat.Walk(); err != nil {
		fmt.Println("Cat are walkable!")
	} else {
		fmt.Println("Cat are not walkable!")
	}
	horse := generated.NewHorse()
	speaker = horse.Speak()
	fmt.Printf("a Horse says: %s\n", speaker)
	if err := horse.Walk(); err != nil {
		fmt.Println("Horse are walkable!")
	} else {
		fmt.Println("Horse are not walkable!")
	}
	parrot := generated.NewParrot()
	speaker = parrot.Speak()
	fmt.Printf("a Parrot says: %s\n", speaker)
	if err := parrot.Walk(); err != nil {
		fmt.Println("Parrot are walkable!")
	} else {
		fmt.Println("Parrot are not walkable!")
	}
	petDog := IdentPet(dog)
	fmt.Printf("pet Dog says: %s\n", petDog.Speak())
	petCat := IdentPet(cat)
	fmt.Printf("pet Cat says: %s\n", petCat.Speak())
	petHorse := IdentPet(horse)
	fmt.Printf("pet Horse says: %s\n", petHorse.Speak())
	petParrot := IdentPet(parrot)
	fmt.Printf("pet Parrot says: %s\n", petParrot.Speak())
}
func IdentPet(p generated.Pet) generated.Pet {
	return p
}
