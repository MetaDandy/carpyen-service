package enum

type Material string

const (
	Forniture Material = "MUEBLE"
	Wood      Material = "MADERA"
	Metal     Material = "METAL"
	Plastic   Material = "PLASTICO"
	Glass     Material = "VIDRIO"
	Paint     Material = "PINTURA"
	Other     Material = "OTRO"
)

func (m Material) IsValid() bool {
	switch m {
	case Forniture, Wood, Metal, Plastic, Glass, Paint, Other:
		return true
	}
	return false
}

func (m Material) String() string {
	return string(m)
}

func MaterialToArray() []string {
	return []string{
		string(Forniture),
		string(Wood),
		string(Metal),
		string(Plastic),
		string(Glass),
		string(Paint),
		string(Other),
	}
}
