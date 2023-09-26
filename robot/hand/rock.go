package hand

type RockHand string

func (r RockHand) Attack() string {
	return "Удар тяжёлым роком! Все живые существа в радиусе 5м оглохли на ближайшие 10 минут"
}
