package fruit

type Fruit int

const (
	Apple Fruit = iota
	Banana
	Cherry
	Orange
	Grape
)

func (i Fruit) String() string {
	switch i {
	case Apple:
		return "Apple"
	case Banana:
		return "Banana"
	case Cherry:
		return "Cherry"
	case Orange:
		return "Orange"
	case Grape:
		return "Grape"
	default:
		return "Unknown"
	}
}
