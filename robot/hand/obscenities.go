package hand

type ObscenitiesHand string

func (o ObscenitiesHand) Attack() string {
	return "Удар отборнейшими выражениями! Все живые и не живые существа в радиусе 5м дезориентированы на ближайшие 10 минут"
}
