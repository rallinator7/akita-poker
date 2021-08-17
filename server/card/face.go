package card

// the value a card can be
type Face int

const (
	Two Face = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func Faces() []Face {
	return []Face{
		Two,
		Three,
		Four,
		Five,
		Six,
		Seven,
		Eight,
		Nine,
		Ten,
		Jack,
		Queen,
		King,
		Ace,
	}
}

// needed to implement Stringer interface
// allows for the string value of the suit to be printed out instead of the int value
func (face Face) String() string {
	faces := []string{
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
		"Ace",
	}

	return faces[int(face)]
}
