package main

import (
	"fmt"
)

type Animal struct {
	id   string
	name string
}
type Animals []Animal

type hoge struct {
	aaa []Animal
}

func main() {
	var dog Animal = Animal{id: "1", name: "dog"}
	var cat Animal = Animal{id: "2", name: "cat"}

	var animals Animals

	animals = append(animals, dog)
	animals = append(animals, cat)
	/*
		var a hoge = hoge{
			aaa: animals,
		}
	*/
	for _, animal := range animals {
		fmt.Println(animal.name)
	}
}
