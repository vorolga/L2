package main

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

import "fmt"

type State interface {
	Do()
}

type Item struct {
	state State
}

func NewItem() *Item {
	return &Item{state: &DefaultState{}}
}

func (i *Item) ChangeState(s State) {
	i.state = s
}

func (i *Item) Do() {
	i.state.Do()
}

type DefaultState struct {
}

func (s *DefaultState) Do() {
	fmt.Println("Default state")
}

type State1 struct {
}

func (s *State1) Do() {
	fmt.Println("State 1")
}

type State2 struct {
}

func (s *State2) Do() {
	fmt.Println("State 2")
}

type State3 struct {
}

func (s *State3) Do() {
	fmt.Println("State 3")
}

func main() {
	i := NewItem()
	i.Do()

	i.ChangeState(&State1{})
	i.Do()

	i.ChangeState(&State2{})
	i.Do()

	i.ChangeState(&State3{})
	i.Do()
}
