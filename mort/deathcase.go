package mort

type DeathCaser interface {
	ReadData(data string)
}

type DeathCase struct {
}

func NewDeathCase(data string) *DeathCase {
	dc := new(DeathCase)
	return dc
}

func (dc *DeathCase) ReadData(data string) {

}
