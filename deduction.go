package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	PercentDeduction = iota
	AbsoluteDeduction
)

type Deduction struct {
	deductionType   int
	deductionAmount float64
}

func (d *Deduction) String() string {
	var deductionType string

	switch d.deductionType {
	case PercentDeduction:
		deductionType = "Percent"
	case AbsoluteDeduction:
		deductionType = "Absolute"
	}

	return fmt.Sprintf("%v deduction of %f", deductionType, d.deductionAmount)
}

func (d *Deduction) Set(value string) error {
	lastCharacter := value[len(value)-1]

	var deductionAmountString string

	if lastCharacter == '%' {
		d.deductionType = PercentDeduction
		deductionAmountString = value[:len(value)-1]
	} else {
		d.deductionType = AbsoluteDeduction
		deductionAmountString = value
	}

	deductionAmount, err := strconv.ParseFloat(deductionAmountString, 64)
	d.deductionAmount = deductionAmount
	return err
}

type DeductionList []Deduction

func (dl *DeductionList) String() string {
	builder := strings.Builder{}
	for _, deduction := range *dl {
		builder.WriteString(deduction.String())
		builder.WriteString(", ")
	}

	return builder.String()
}

func (dl *DeductionList) Set(value string) error {
	if len(*dl) > 0 {
		return errors.New("DeductionList already set")
	}

	for _, deductionString := range strings.Split(value, ",") {
		var deduction Deduction
		err := deduction.Set(deductionString)

		if err != nil {
			return err
		}

		*dl = append(*dl, deduction)
	}
	return nil
}

func calculateDeductedIncome(income float64, deductionList DeductionList) float64 {
	totalDeductions := 0.0
	remainingIncome := income

	for _, deduction := range deductionList {
		switch deduction.deductionType {
		case PercentDeduction:
			deduction := remainingIncome * deduction.deductionAmount / 100
			totalDeductions += deduction
			remainingIncome -= deduction
		case AbsoluteDeduction:
			deduction := deduction.deductionAmount
			totalDeductions += deduction
			remainingIncome -= deduction

		default:
			panic(fmt.Sprintf("Unknown deduction type: %v", deduction.deductionType))
		}
	}

	return totalDeductions
}
