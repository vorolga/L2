package main

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

import "fmt"

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type command interface {
	execute()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type device interface {
	on()
	off()
}

type notebook struct {
	isRunning bool
}

func (t *notebook) on() {
	t.isRunning = true
	fmt.Println("Turning notebook on")
}

func (t *notebook) off() {
	t.isRunning = false
	fmt.Println("Turning notebook off")
}

func main() {
	notebook := &notebook{}
	onCommand := &onCommand{
		device: notebook,
	}
	offCommand := &offCommand{
		device: notebook,
	}
	onButton := &button{
		command: onCommand,
	}
	onButton.press()
	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
