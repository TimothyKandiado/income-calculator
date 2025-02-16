package tax

import "math"

type Threshold struct {
	limit         float64
	taxPercentage float64
}

type Strategy struct {
	thresholds []Threshold
}

func (t Strategy) CalculateTax(grossAmount float64) float64 {
	taxedSum := 0.0
	remainingAmount := grossAmount

	for _, threshold := range t.thresholds {
		var taxableAmount float64

		if remainingAmount <= 0 {
			break
		}

		if remainingAmount >= threshold.limit {
			taxableAmount = threshold.limit
			remainingAmount -= threshold.limit
		} else {
			taxableAmount = remainingAmount
			remainingAmount = 0
		}

		taxedSum += taxableAmount * threshold.taxPercentage / 100
	}

	return taxedSum
}

func NewMalawianTaxStrategy() Strategy {
	thresholds := []Threshold{
		Threshold{
			limit:         150_000,
			taxPercentage: 0.0,
		},
		Threshold{
			limit:         350_000,
			taxPercentage: 25.0,
		},
		Threshold{
			limit:         2_050_000,
			taxPercentage: 30.0,
		},
		Threshold{
			limit:         math.Inf(1),
			taxPercentage: 35.0,
		},
	}

	return Strategy{
		thresholds: thresholds,
	}
}
