package hand

type BubbleHand string

func (b BubbleHand) Attack() string {
	return "Удар потоком мыльных пузырей! Все живые существа в радиусе 5м ослепли на ближайшие 10 минут"
}
