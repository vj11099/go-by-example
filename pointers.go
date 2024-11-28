package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial value: ", i)
	fmt.Println()

	zeroVal(i)
	fmt.Println("After updating value in zeroVal func: ", i)
	fmt.Println()

	zeroPtr(&i)
	fmt.Println("After pointing i's address to value 0: ", i)
	fmt.Println()

	fmt.Println("Address/Pointer of i", &i)
	fmt.Println("")
	pointers()
}

//additional pointers practice

type animal struct {
	species string
	name    string
	wild    bool
}

func pointers() {
	i := 14
	a := &i
	*a = 12

	newAnimal := &animal{species: "Dog", wild: true}
	fmt.Println("New Animal: ", newAnimal, "\n")
	// fmt.Printf("New Animal's address:%p \n", &newAnimal)

	pet := &newAnimal

	// fmt.Println("Pet Animal: ", **pet)
	// fmt.Println("Pet Animal's pointing to: ", pet)
	// fmt.Printf("Pet Animal's Address: %p \n", &pet)

	adoptAnimal(*pet, "Max")

	fmt.Println((*newAnimal).name, "is my", (*newAnimal).species)
	// *&pet.wild = false
	fmt.Println("")
}

func adoptAnimal(animal *animal, name string) {
	(*animal).name = name
	(*animal).wild = false
}
