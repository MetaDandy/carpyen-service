package enum

type Measure string

const (
	Meter       Measure = "METRO"
	SquareMeter Measure = "METRO_CUADRADO"
	MeasureUnit Measure = "UNIDAD"
	Liter       Measure = "LITRO"
	Kilogram    Measure = "KILOGRAMO"
	Pound       Measure = "LIBRA"
	CubicMeter  Measure = "METRO_CUBICO"
	CubicLiter  Measure = "LITRO_CUBICO"
)

func (m Measure) IsValid() bool {
	switch m {
	case Meter, SquareMeter, MeasureUnit, Liter, Kilogram, Pound, CubicMeter, CubicLiter:
		return true
	}
	return false
}

func (m Measure) String() string {
	return string(m)
}
