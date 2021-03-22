package main

import (
	"fmt"
	"reflect"
)

// struct, value type
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string `required max:"100"` // tag
	Origin string
}

type Bird struct {
	Animal   // embedded a Animal struct
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// Map
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 21345,
		"Texas":      12345,
		"New York":   11345,
		"Ohio":       37489,
	}
	statePopulations["Georgia"] = 10567
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["New York"])
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	fmt.Println(len(statePopulations))

	// to check if the key is in the map, use "comma, ok syntax"
	_, ok := statePopulations["Ohio"]
	fmt.Println(ok)

	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)
	// Ohio doesn't exist in either of the map since map is passed by reference

	// Struct
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])

	aAnonymousDocStruct := struct{ name string }{name: "John Pertwee"}
	fmt.Println(aAnonymousDocStruct)
	anotherDoc := aAnonymousDocStruct
	// anotherDoc := &aAnonymousDocStruct

	anotherDoc.name = "Tom Baker"
	fmt.Println(aAnonymousDocStruct) // {John Pertwee}
	fmt.Println(anotherDoc)          // {Tom Baker}

	// embedded, composite
	bird := Bird{}
	bird.Name = "Emu"
	bird.Origin = "Austrilia"
	bird.SpeedKPH = 48
	bird.CanFly = false
	fmt.Println(bird)

	// tag
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
