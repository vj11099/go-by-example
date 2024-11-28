package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      int
}

func newPerson(name, lastName string) *person { //returns a pointer
	return &person{name: name, lastName: lastName, age: 42}
}

func main() {

	fmt.Println(person{"Bob", "Doe", 20}, "\n")

	fmt.Println(person{lastName: "Williams", name: "Alice", age: 20}, "\n")

	fmt.Println(person{name: "Francis"}, "\n")

	p := person{name: "John", lastName: "Parker", age: 25}
	// Accessing memory address of specific fields
	fmt.Printf("Memory address of  person: %p", &p) // Address of the 'person'
	fmt.Println()
	fmt.Println("Memory address of name:", &p.name)         // Address of the 'name' field
	fmt.Println("Memory address of lastName:", &p.lastName) // Address of the 'lastName' field
	fmt.Println("Memory address of age:", &p.age)           // Address of the 'age' field
	fmt.Println()

	fmt.Println(&person{name: "Paul", lastName: "Banks", age: 27}, "\n")

	//using a function to initiate a struct
	fmt.Println(*newPerson("Josh", "Eubanks"), "\n") //de-referencing the pointer to to print the struct fields

	s := person{name: "Sean", lastName: "Doe", age: 41}
	fmt.Println("s.name:", s.name)

	sp := &s
	fmt.Println("sp.name:", sp.name)

	sp.name = "Shaun" //mutating a struct
	fmt.Println("new s.name = sp.name:", s.name, "\n")

	dog := struct { // inline struct
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog.name)

}

// (**pet).name = "Max"
// fmt.Println("New Animal:", newAnimal)
// fmt.Println("Pet:", *pet)
