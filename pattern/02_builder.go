package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	if builderType == "stone" {
		return &stoneBuilder{}
	}
	if builderType == "wood" {
		return &woodBuilder{}
	}
	return nil
}

type stoneBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newStoneBuilder() *stoneBuilder {
	return &stoneBuilder{}
}

func (b *stoneBuilder) setWindowType() {
	b.windowType = "Stone Window"
}

func (b *stoneBuilder) setDoorType() {
	b.doorType = "Stone Door"
}

func (b *stoneBuilder) setNumFloor() {
	b.floor = 3
}

func (b *stoneBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type woodBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newWoodBuilder() *woodBuilder {
	return &woodBuilder{}
}

func (b *woodBuilder) setWindowType() {
	b.windowType = "Wood Window"
}

func (b *woodBuilder) setDoorType() {
	b.doorType = "Wood Door"
}

func (b *woodBuilder) setNumFloor() {
	b.floor = 2
}

func (b *woodBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type house struct {
	windowType string
	doorType   string
	floor      int
}

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func main() {
	stoneBuilder := getBuilder("stone")
	woodBuilder := getBuilder("wood")

	director := newDirector(stoneBuilder)
	stoneHouse := director.buildHouse()

	fmt.Printf("stone House Door Type: %s\n", stoneHouse.doorType)
	fmt.Printf("stone House Window Type: %s\n", stoneHouse.windowType)
	fmt.Printf("stone House Num Floor: %d\n", stoneHouse.floor)

	director.setBuilder(woodBuilder)
	woodHouse := director.buildHouse()

	fmt.Printf("\nwood House Door Type: %s\n", woodHouse.doorType)
	fmt.Printf("wood House Window Type: %s\n", woodHouse.windowType)
	fmt.Printf("wood House Num Floor: %d\n", woodHouse.floor)
}
