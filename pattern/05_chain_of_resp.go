package main

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

import "fmt"

type action interface {
	execute(fridge2 *fridge)
	setNext(action)
}

type openFridge struct {
	next action
}

func (o *openFridge) execute(f *fridge) {
	if f.opened {
		fmt.Println("fridge already opened")
		o.next.execute(f)
		return
	}
	fmt.Println("opening fridge")
	f.opened = true
	o.next.execute(f)
}

func (o *openFridge) setNext(next action) {
	o.next = next
}

type takeSausage struct {
	next action
}

func (t *takeSausage) execute(f *fridge) {
	if f.sausageIsTaken {
		fmt.Println("sausage already taken")
		t.next.execute(f)
		return
	}
	fmt.Println("taking sausage")
	f.sausageIsTaken = true
	t.next.execute(f)
}

func (d *takeSausage) setNext(next action) {
	d.next = next
}

type closeFridge struct {
	next action
}

func (c *closeFridge) execute(f *fridge) {
	if f.closed {
		fmt.Println("fridge already closed")
		return
	}
	fmt.Println("closing fridge")
	f.closed = true
}

func (c *closeFridge) setNext(a action) {
	c.next = a
}

type fridge struct {
	brand          string
	opened         bool
	sausageIsTaken bool
	closed         bool
}

func main() {
	closeFridge := &closeFridge{}

	takeSausage := &takeSausage{}
	takeSausage.setNext(closeFridge)

	openFridge := &openFridge{}
	openFridge.setNext(takeSausage)

	fridge := &fridge{brand: "cold"}
	openFridge.execute(fridge)
}
