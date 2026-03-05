package enum

type Unit string

const (
	Linear Unit = "LINEAL"
	Square Unit = "CUADRADO"
	Piece  Unit = "PIEZA"
)

func (u Unit) IsValid() bool {
	switch u {
	case Linear, Square, Piece:
		return true
	}
	return false
}

func (u Unit) String() string {
	return string(u)
}
func UnitToArray() []string {
	return []string{
		string(Linear),
		string(Square),
		string(Piece),
	}
}
