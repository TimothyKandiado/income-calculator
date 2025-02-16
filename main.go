package main

import (
	"flag"
	"fmt"
	"income-calculator/tax"
)

func main() {
	var income float64
	var currency string
	var deductions DeductionList

	const (
		defaultIncome   = 0.0
		defaultCurrency = "MWK"

		incomeUsage   = "gross income"
		currencyUsage = "currency"

		deductionsUsage = "Comma-separated list of deductions"
	)

	flag.Float64Var(&income, "income", defaultIncome, incomeUsage)
	flag.Float64Var(&income, "i", 0.0, incomeUsage+"(shorthand)")

	flag.StringVar(&currency, "currency", defaultCurrency, currencyUsage)
	flag.StringVar(&currency, "c", defaultCurrency, currencyUsage+"(shorthand)")

	flag.Var(&deductions, "deductions", deductionsUsage)
	flag.Var(&deductions, "d", deductionsUsage+"(shorthand)")

	flag.Parse()

	calculateIncome(income, currency, deductions)
}

func calculateIncome(grossIncome float64, currency string, deductions DeductionList) {
	taxStrategy := tax.NewMalawianTaxStrategy()
	taxedIncome := taxStrategy.CalculateTax(grossIncome)
	remainingIncome := grossIncome - taxedIncome

	deductedIncome := calculateDeductedIncome(remainingIncome, deductions)

	remainingIncome -= deductedIncome

	fmt.Printf("Taxed Income: 		%v %6.2f\n", currency, taxedIncome)
	fmt.Printf("Deducted Income: 	%v %6.2f\n", currency, deductedIncome)
	fmt.Printf("Remaining Income: 	%v %6.2f\n", currency, remainingIncome)
}
