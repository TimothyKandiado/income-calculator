package interfaces

type ITaxStrategy interface {
	CalculateTax(grossAmount float64) float64
}
