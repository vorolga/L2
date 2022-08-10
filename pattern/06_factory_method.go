package main

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

import "fmt"

type iMusicalInstrument interface {
	setName(name string)
	setStrings(power int)
	getName() string
	getStrings() int
}

type musicalInstrument struct {
	name    string
	strings int
}

func (g *musicalInstrument) setName(name string) {
	g.name = name
}

func (g *musicalInstrument) getName() string {
	return g.name
}

func (g *musicalInstrument) setStrings(power int) {
	g.strings = power
}

func (g *musicalInstrument) getStrings() int {
	return g.strings
}

type guitar struct {
	musicalInstrument
}

func newGuitar() iMusicalInstrument {
	return &guitar{
		musicalInstrument: musicalInstrument{
			name:    "guitar",
			strings: 6,
		},
	}
}

type ukulele struct {
	musicalInstrument
}

func newUkulele() iMusicalInstrument {
	return &ukulele{
		musicalInstrument: musicalInstrument{
			name:    "ukulele",
			strings: 4,
		},
	}
}

func getMusicalInstrument(musicalInstrumentType string) (iMusicalInstrument, error) {
	if musicalInstrumentType == "guitar" {
		return newGuitar(), nil
	}
	if musicalInstrumentType == "ukulele" {
		return newUkulele(), nil
	}
	return nil, fmt.Errorf("Wrong musical instrument type passed")
}

func main() {
	guitar, _ := getMusicalInstrument("guitar")
	ukulele, _ := getMusicalInstrument("ukulele")
	printDetails(guitar)
	printDetails(ukulele)
}

func printDetails(i iMusicalInstrument) {
	fmt.Printf("Musical Instrument: %s", i.getName())
	fmt.Println()
	fmt.Printf("Strings: %d", i.getStrings())
	fmt.Println()
}
