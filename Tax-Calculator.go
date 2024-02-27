/*********************
Name:		DELA CRUZ, Frances Julianne
			GAMBOA, Mikkel Dominic
			VERANO, Carl Matthew
			YU, Hanz Patrick
Language:	Go
Paradigm:	MultiParadigm: Imperative, Concurrent, Object-Oriented
********************/

package main

import "fmt"

// PHP represents US dollar amount in terms of cents
type PHP int64

// ToPHP converts a float64 to PHP
// e.g. 1.23 to $1.23, 1.345 to $1.35
func ToPHP(f float64) PHP {
	return PHP((f * 100) + 0.5)
}

// Float64 converts a PHP to float64
func (m PHP) Float64() float64 {
	x := float64(m)
	x = x / 100
	return x
}

// Multiply safely multiplies a PHP value by a float64, rounding
// to the nearest cent.
func (m PHP) Multiply(f float64) PHP {
	x := (float64(m) * f) + 0.5
	return PHP(x)
}

// String returns a formatted PHP value
func (m PHP) String() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("%.2f", x)
}

func getPhilHealth(monthlyIncome float64) float64 {

	var philHealth float64

	if monthlyIncome <= 10000 {
		philHealth = 400.00
	} else if monthlyIncome >= 80000 {
		philHealth = 3200.00
	} else {
		philHealthPHP := ToPHP(monthlyIncome)
		philHealth = philHealthPHP.Multiply(0.04).Float64()
	}

	return philHealth
}

func getPagibig(monthlyIncome float64) float64 {

	var pagibig float64
	var maxContribution float64

	if monthlyIncome > 5000 {
		maxContribution = 5000
	} else {
		maxContribution = monthlyIncome
	}

	maxContributionPHP := ToPHP(maxContribution)

	if maxContribution > 1500 {
		pagibig = maxContributionPHP.Multiply(0.02).Float64()
	} else {
		pagibig = maxContributionPHP.Multiply(0.01).Float64()
	}

	return pagibig
}

func getTotalContributions(SSS float64,
	PhilHealth float64,
	pagIBIG float64) float64 {

	SSSPHP := ToPHP(SSS)
	PhilHealthPHP := ToPHP(PhilHealth)
	pagIBIGPHP := ToPHP(pagIBIG)

	return (SSSPHP + PhilHealthPHP + pagIBIGPHP).Float64()
}

func getEmployeeSSS(monthlyIncome float64) float64 {
	type SSS struct {
		employeeSSS         float64
		MPF                 float64
		monthlySalaryCredit float64
	}
	sss := SSS{employeeSSS: 0, MPF: 0, monthlySalaryCredit: 0}

	if monthlyIncome < 1000 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 0
	}
	if monthlyIncome >= 1000 && monthlyIncome < 3250 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 3000
	}
	if monthlyIncome >= 3250 && monthlyIncome < 3750 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 3500
	}
	if monthlyIncome >= 3750 && monthlyIncome < 4250 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 4000
	}
	if monthlyIncome >= 4250 && monthlyIncome <= 4749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 4500
	}
	if monthlyIncome >= 4750 && monthlyIncome <= 5249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 5000
	}
	if monthlyIncome >= 5250 && monthlyIncome <= 5749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 5500
	}
	if monthlyIncome >= 5750 && monthlyIncome <= 6249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 6000
	}
	if monthlyIncome >= 6250 && monthlyIncome <= 6749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 6500
	}
	if monthlyIncome >= 6750 && monthlyIncome <= 7249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 7000
	}
	if monthlyIncome >= 7250 && monthlyIncome <= 7749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 7500
	}
	if monthlyIncome >= 7750 && monthlyIncome <= 8249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 8000
	}
	if monthlyIncome >= 8250 && monthlyIncome <= 8749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 8500
	}
	if monthlyIncome >= 8750 && monthlyIncome <= 9249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 9000
	}
	if monthlyIncome >= 9250 && monthlyIncome <= 9749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 9500
	}
	if monthlyIncome >= 9750 && monthlyIncome <= 10249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 10000
	}
	if monthlyIncome >= 10250 && monthlyIncome <= 10749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 10500
	}
	if monthlyIncome >= 10750 && monthlyIncome <= 11249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 11000
	}
	if monthlyIncome >= 11250 && monthlyIncome <= 11749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 11500
	}
	if monthlyIncome >= 11750 && monthlyIncome <= 12249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 12000
	}
	if monthlyIncome >= 12250 && monthlyIncome <= 12749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 12500
	}
	if monthlyIncome >= 12750 && monthlyIncome <= 13249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 13000
	}
	if monthlyIncome >= 13250 && monthlyIncome <= 13749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 13500
	}
	if monthlyIncome >= 13750 && monthlyIncome <= 14249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 14000
	}
	if monthlyIncome >= 14250 && monthlyIncome <= 14749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 14500
	}
	if monthlyIncome >= 14750 && monthlyIncome <= 15249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 15000
	}
	if monthlyIncome >= 15250 && monthlyIncome <= 15749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 15500
	}
	if monthlyIncome >= 15750 && monthlyIncome <= 16249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 16000
	}
	if monthlyIncome >= 16250 && monthlyIncome <= 16749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 16500
	}
	if monthlyIncome >= 16750 && monthlyIncome <= 17249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 17000
	}
	if monthlyIncome >= 17250 && monthlyIncome <= 17749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 17500
	}
	if monthlyIncome >= 17750 && monthlyIncome <= 18249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 18000
	}
	if monthlyIncome >= 18250 && monthlyIncome <= 18749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 18500
	}
	if monthlyIncome >= 18750 && monthlyIncome <= 19249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 19000
	}
	if monthlyIncome >= 19250 && monthlyIncome <= 19749.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 19500
	}
	if monthlyIncome >= 19750 && monthlyIncome <= 20249.99 {
		sss.MPF = 0
		sss.monthlySalaryCredit = 20000
	}
	if monthlyIncome >= 20250 && monthlyIncome <= 20749.99 {
		sss.MPF = 500
		sss.monthlySalaryCredit = 20500
	}
	if monthlyIncome >= 20750 && monthlyIncome <= 21249.99 {
		sss.MPF = 1000
		sss.monthlySalaryCredit = 21000
	}
	if monthlyIncome >= 21250 && monthlyIncome <= 21749.99 {
		sss.MPF = 1500
		sss.monthlySalaryCredit = 21500
	}
	if monthlyIncome >= 21750 && monthlyIncome <= 22249.99 {
		sss.MPF = 2000
		sss.monthlySalaryCredit = 22000
	}
	if monthlyIncome >= 22250 && monthlyIncome <= 22749.99 {
		sss.MPF = 2500
		sss.monthlySalaryCredit = 22500
	}
	if monthlyIncome >= 22750 && monthlyIncome <= 23249.99 {
		sss.MPF = 3000
		sss.monthlySalaryCredit = 23000
	}
	if monthlyIncome >= 23250 && monthlyIncome <= 23749.99 {
		sss.MPF = 3500
		sss.monthlySalaryCredit = 23500
	}
	if monthlyIncome >= 23750 && monthlyIncome <= 24249.99 {
		sss.MPF = 4000
		sss.monthlySalaryCredit = 24000
	}
	if monthlyIncome >= 24250 && monthlyIncome <= 24749.99 {
		sss.MPF = 4500
		sss.monthlySalaryCredit = 24500
	}
	if monthlyIncome >= 24750 {
		sss.MPF = 5000
		sss.monthlySalaryCredit = 25000
	}

	sss.employeeSSS = ((sss.monthlySalaryCredit - sss.MPF) * 0.045) + (sss.MPF * 0.045)

	return sss.employeeSSS
}

func getIncomeTax(monthlyIncome float64, totalContributions float64) float64 {

	var taxableIncome float64 = monthlyIncome - totalContributions
	var pOverCL float64
	var incomeTax float64

	if taxableIncome < 20833 {
		incomeTax = 0
	} else if taxableIncome >= 20833 && taxableIncome < 33333 {
		pOverCL = (taxableIncome - 20833) * 0.2
		incomeTax = 0.00 + pOverCL
	} else if taxableIncome >= 33333 && taxableIncome < 66667 {
		pOverCL = (taxableIncome - 33333) * 0.25
		incomeTax = 2500 + pOverCL
	} else if taxableIncome >= 66667 && taxableIncome < 166667 {
		pOverCL = (taxableIncome - 66667) * 0.3
		incomeTax = 10833.33 + pOverCL
	} else if taxableIncome >= 166667 && taxableIncome < 666667 {
		pOverCL = (taxableIncome - 166667) * 0.32
		incomeTax = 40833.33 + pOverCL
	} else if taxableIncome >= 666667 {
		pOverCL = (taxableIncome - 666667) * 0.35
		incomeTax = 200833.33 + pOverCL
	}

	return incomeTax
}

func main() {
	var monthlyIncome float64
	var incomeTax float64
	var pagIBIG float64
	var SSS float64
	var PhilHealth float64
	var totalContributions float64

	fmt.Print("Input Monthly Income: ")
	fmt.Scan(&monthlyIncome)

	// ASSIGNING OF FUNCTIONS TO VALUES
	SSS = getEmployeeSSS(monthlyIncome)
	PhilHealth = getPhilHealth(monthlyIncome)
	pagIBIG = getPagibig(monthlyIncome)
	totalContributions = getTotalContributions(SSS, PhilHealth, pagIBIG)
	incomeTax = getIncomeTax(monthlyIncome, totalContributions)

	//PHP Computations
	monthlyIncomePHP := ToPHP(monthlyIncome)
	incomeTaxPHP := ToPHP(incomeTax)
	totalContributionsPHP := ToPHP(totalContributions)

	// DISPLAY OUTPUTS
	fmt.Println("Monthly Income:", monthlyIncomePHP.String())
	fmt.Println()
	fmt.Println("TAX COMPUTATION")
	fmt.Println("Income Tax:", incomeTaxPHP.String())
	fmt.Println("Net Pay after Tax:", (monthlyIncomePHP - incomeTaxPHP).String())
	fmt.Println()
	fmt.Println("MONTHLY CONTRIBUTIONS")
	fmt.Println("SSS:", ToPHP(SSS).String())
	fmt.Println("PhilHealth:", ToPHP(PhilHealth).String())
	fmt.Println("Pag-IBIG: ", ToPHP(pagIBIG).String())
	fmt.Println("Total Contributions: ", totalContributionsPHP.String())
	fmt.Println()
	fmt.Println("Total Deductions: ", (incomeTaxPHP + totalContributionsPHP).String())
	fmt.Println("Net Pay after Deductions:", (monthlyIncomePHP - (incomeTaxPHP + totalContributionsPHP)).String())
}
