package tax

import (
	"fmt"
	"income-calculator/interfaces"
	"testing"
)

func Test_Malawian_Tax_Strategy(t *testing.T) {
	type TestCase struct {
		grossIncome             float64
		expectedRemainingIncome float64
	}

	strategy := NewMalawianTaxStrategy()

	tests := []TestCase{
		{grossIncome: 150_000, expectedRemainingIncome: 150_000},
		{grossIncome: 500_000, expectedRemainingIncome: 412_500},
		{grossIncome: 1_000_000, expectedRemainingIncome: 762_500},
	}

	for _, test := range tests {
		err := testRemainingIncome(&strategy, test.grossIncome, test.expectedRemainingIncome)
		if err != nil {
			t.Fatalf("%v", err.Error())
		}
	}
}

func testRemainingIncome(strategy interfaces.ITaxStrategy, grossIncome float64, expectedRemainingIncome float64) error {
	taxedIncome := strategy.CalculateTax(grossIncome)
	remainingIncome := grossIncome - taxedIncome

	if remainingIncome != expectedRemainingIncome {
		return fmt.Errorf("for gross income = %f, remaining income is expected to be %f, but found %f instead", grossIncome, expectedRemainingIncome, remainingIncome)
	}
	return nil
}
