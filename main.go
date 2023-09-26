package main

import (
	"github.com/KseniiaSalmina/InterfaceRobots/robot"
	"github.com/KseniiaSalmina/InterfaceRobots/robot/hand"
	"github.com/KseniiaSalmina/InterfaceRobots/robot/leg"
)

func main() {
	Vasiliy := robot.NewRobot("Vasiliy", hand.ObscenitiesHand("hand"), leg.RunLeg("leg"))
	Petr := robot.NewRobot("Petr", hand.RockHand("hand"), leg.FlyLeg("leg"))
	Eduard := robot.NewRobot("Eduard", hand.BubbleHand("hand"), leg.FlyLeg("leg"))

	Vasiliy.Test()
	Petr.Test()
	Eduard.Test()
}
