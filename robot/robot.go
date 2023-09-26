package robot

import "fmt"

type Shooter interface {
	Attack() string
}

type Mover interface {
	Move() string
}

type Robot struct {
	Name string
	Hand Shooter
	Leg  Mover
}

func NewRobot(name string, hand Shooter, leg Mover) Robot {
	return Robot{
		Name: name,
		Hand: hand,
		Leg:  leg,
	}
}

func (r Robot) Test() {
	fmt.Printf("Робот %s запущен\n", r.Name)

	fmt.Printf("...................\nВнимание! %s активирует оружие!\n%s\n", r.Name, r.Hand.Attack())
	fmt.Printf("...................\n%s дает по съебам!\n%s\n\n", r.Name, r.Leg.Move())
}
