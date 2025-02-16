# Income Calculator
a cli tool for automatically calculating income tax and other deductions

## Usage
income-calculator [options]

e.g:

income-calculator --help

income-calculator -income 500000 -currency MWK -deductions 5%,5000
income-calculator -i 1_000_000 -c MWK -d 5_000,5%

### Income
Supports any float64 value

### Currency
A currency string to use, "MWK" is the default

### Deductions
Requires a comma delimited list of deductions

The program supports too types of deductions: (percentage deductions and absolute deductions)

5% : denotes a percentage deduction

5000 : denotes an absolute deduction

--d 5%,10%,5000 means that deduct 5%, then 10%, then 5000 from the income after tax