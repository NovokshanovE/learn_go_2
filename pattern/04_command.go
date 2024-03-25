package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

import "fmt"

// Интерфейс команды
type Command interface {
	Execute()
}

// Конкретная команда
type SwitchOnCommand struct {
	device Device
}

func (c *SwitchOnCommand) Execute() {
	c.device.TurnOn()
}

type SwitchOffCommand struct {
	device Device
}

func (c *SwitchOffCommand) Execute() {
	c.device.TurnOff()
}

// Получатель команды
type Device struct {
	name string
}

func (d *Device) TurnOn() {
	fmt.Printf("%s включено\n", d.name)
}

func (d *Device) TurnOff() {
	fmt.Printf("%s выключено\n", d.name)
}

// Инициатор команды
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func CommandPattern() {
	tv := Device{name: "Телевизор"}
	switchOn := &SwitchOnCommand{device: tv}
	switchOff := &SwitchOffCommand{device: tv}

	remote := &RemoteControl{command: switchOn}
	remote.PressButton()

	remote.command = switchOff
	remote.PressButton()
}
