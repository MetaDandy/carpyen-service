package enum

type Product string

const (
	Chair        Product = "SILLA"
	Table        Product = "MESA"
	Sofa         Product = "SOFA"
	Bed          Product = "CAMA"
	Cabinet      Product = "GABINETE"
	Desk         Product = "ESCRITORIO"
	Shelf        Product = "ESTANTE"
	Lamp         Product = "LAMPARA"
	Rug          Product = "ALFOMBRA"
	Curtain      Product = "CORTINA"
	OtherProduct Product = "OTRO"
)

func (p Product) IsValid() bool {
	switch p {
	case Chair, Table, Sofa, Bed, Cabinet, Desk, Shelf, Lamp, Rug, Curtain, OtherProduct:
		return true
	}
	return false
}

func (p Product) String() string {
	return string(p)
}
