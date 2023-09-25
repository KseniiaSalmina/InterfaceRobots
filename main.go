package main

import (
	"fmt"
)

func main() {
	fmt.Println("Старт программы")
	start()
	fmt.Println("Конец программы")
}

type bubbleHand string
type rockHand string
type obscenitiesHand string

func (bubbleHand) attack() {
	fmt.Println("Удар потоком мыльных пузырей! Все живые существа в радиусе 5м ослепли на ближайшие 10 минут")
}

func (rockHand) attack() {
	fmt.Println("Удар тяжёлым роком! Все живые существа в радиусе 5м оглохли на ближайшие 10 минут")
}

func (obscenitiesHand) attack() {
	fmt.Println("Удар отборнейшими выражениями! Все живые и не живые существа в радиусе 5м дезориентированы на ближайшие 10 минут")
}

type runLeg string
type flyLeg string

func (runLeg) move() {
	fmt.Println("Бег активирован. Робот скрылся в Саратове")
}

func (flyLeg) move() {
	fmt.Println("Взлёт активирован. Робот отправился на Мальдивы")
}

type transformer interface {
	attack()
	move()
}

type Vasiliy struct {
	runLeg
	obscenitiesHand
}

type Petr struct {
	flyLeg
	rockHand
}

type Eduard struct {
	flyLeg
	bubbleHand
}

func start() {
	fmt.Println("Тестирование 1: Василий")
	testRobot(Vasiliy{})
	fmt.Println("Тестирование 2: Пётр")
	testRobot(Petr{})
	fmt.Println("Тестирование 3: Эдуард")
	testRobot(Eduard{})
}

func testRobot(t transformer) {
	t.attack()
	t.move()
}
