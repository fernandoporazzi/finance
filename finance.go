package finance

import (
	"math"
)

// PaymentType is used to express if the the payments are monthy or yearly
type PaymentType int

const (
	Years PaymentType = iota
	Months
)

// PV - Present Value
//
// Present Value (PV) is the current worth of a future sum of money or stream of cash flows given a specified rate of return.
func PV(rate, cashFlow, period float64) float64 {
	rate = rate / 100
	pv := cashFlow / math.Pow(1+rate, period)
	return math.Round(pv*100) / 100
}

// FV - Future Value
//
// Future Value (FV) is the value of an asset or cash at a specified date in the future that is equivalent in value to a specified sum today.
func FV(rate, cashFlow, period float64) float64 {
	rate = rate / 100
	fv := cashFlow * math.Pow((1+rate), period)
	return math.Round(fv*100) / 100
}

// NPV - Net Present Value
//
// Net Present Value (NPV) compares the money received in the future to an amount of money received today, while accounting for time and interest [through the discount rate]. It's based on the principal of time value of money (TVM), which explains how time affects monetary value.
func NPV(rate, initialInvestment float64, cashFlows []float64) float64 {
	rate = rate / 100
	npv := initialInvestment

	for i := 0; i < len(cashFlows); i++ {
		ii := float64(i + 2)
		npv = npv + (cashFlows[i] / math.Pow(1+rate, ii-1))
	}

	return math.Round(npv*100) / 100
}

// ROI - Return On Investment
//
// Return on Investment (ROI) is a simple calculation that tells you the bottom line return of any investment.
func ROI(initialInvestment, earnings float64) float64 {
	roi := (earnings - math.Abs(initialInvestment)) / math.Abs(initialInvestment) * 100
	return math.Round(roi*100) / 100
}

// PP - Payback Period
//
// Payback Period (PP) is the length of time required to recover the cost of an investment.
//
// number of periods takes a 0 value for even cash flows;
// for uneven cash flows, number of periods takes any number of projected periods.
// [cashFlows] takes any number of projected cash flows.
func PP(numberOfPeriods int, initialInvestment float64, cashFlows []float64) float64 {
	// for even cash flows
	if numberOfPeriods == 0 {
		return math.Abs(initialInvestment / cashFlows[0])
	}

	cumulativeCashFlow := initialInvestment
	var yearsCounter float64 = 1

	for i := 0; i < len(cashFlows); i++ {
		cumulativeCashFlow += cashFlows[i]

		if cumulativeCashFlow > 0 {
			yearsCounter += (cumulativeCashFlow - cashFlows[i]) / cashFlows[i]
			break
		} else {
			yearsCounter++
		}
	}

	return yearsCounter
}

func buildNumerator(ratePerPeriod, numInterestAccruals float64, payAtBeginning bool) float64 {
	if payAtBeginning {
		// if payments are made in the beginning of the period, then interest shouldn't be calculated for first period
		numInterestAccruals--
	}

	return ratePerPeriod * math.Pow(1+ratePerPeriod, numInterestAccruals)
}

// AM - Amortization
//
// Amortization is the paying off of debt with a fixed repayment schedule in regular installments over a period of time.
//
// paymentType can be either 'months' or 'years'
func AM(principal, rate, period float64, paymentType PaymentType, payAtBeginning bool) float64 {
	if paymentType != Years && paymentType != Months {
		panic("paymentType should be either Months or Years")
	}

	var numerator, denominator, am float64

	ratePerPeriod := rate / 12 / 100

	if paymentType == Years {
		numerator = buildNumerator(ratePerPeriod, period*12, payAtBeginning)
		denominator = math.Pow(1+ratePerPeriod, period*12) - 1
	} else {
		numerator = buildNumerator(ratePerPeriod, period, payAtBeginning)
		denominator = math.Pow(1+ratePerPeriod, period) - 1
	}

	am = principal * (numerator / denominator)
	return math.Round(am*100) / 100
}

// DF - Discount Factor
//
// The Discount Factor (DF) is the factor by which a future cash flow must be multiplied in order to obtain the present value.
func DF(rate float64, numOfPeriods int) []float64 {
	dfs := make([]float64, numOfPeriods-1)
	var discountFactor float64

	for i := 1; i < numOfPeriods; i++ {
		index := float64(i)
		discountFactor = 1 / math.Pow((1+rate/100), (index-1))
		dfs[i-1] = math.Ceil(discountFactor*1000) / 1000
	}

	return dfs
}

// CI - Compound Interest
//
// Compound Interest is the interest calculated on the initial principal and also on the accumulated interest of previous periods of a deposit or loan.
func CI(rate, numOfCompoundings, principal, numOfPeriods float64) float64 {
	ci := principal * math.Pow((1+((rate/100)/numOfCompoundings)), numOfCompoundings*numOfPeriods)
	return math.Round(ci*100) / 100
}

// CAGR - Compound Annual Growth Rate
//
// Compound Annual Growth Rate (CAGR) is the year-over-year growth rate of an investment over a specified period of time.
func CAGR(beginningValue, endingValue, numOfPeriods float64) float64 {
	cagr := math.Pow((endingValue/beginningValue), 1/numOfPeriods) - 1
	return math.Round(cagr*10000) / 100
}

// LR - Leverage Ratio
//
// Leverage Ratio (LR) is used to calculate the financial leverage of a company or individual to get an idea of the methods of financing or to measure ability to meet financial obligations.
func LR(totalLiabilities, totalDebts, totalIncome float64) float64 {
	return (totalLiabilities + totalDebts) / totalIncome
}

// R72 - Rule of 72
//
// Rule of 72 (R72) is a rule stating that in order to find the number of years required to double your money at a given interest rate, you divide the compound return into 72.
func R72(rate float64) float64 {
	return 72 / rate
}

// PMT - Payment
//
// PMT or annuity payment is an inflow or outflow amount that occurs at each compounding period of a financial stream.
func PMT(rate, numOfPayments, principal float64) float64 {
	rate = rate / 100
	pmt := -(principal * rate) / (1 - math.Pow(1+rate, -numOfPayments))
	return math.Round(pmt*100) / 100
}

// IAR calculates the Inflation-adjusted return
//
// Measure the return taking into account the time period's inflation rate
func IAR(investmentReturn, inflationRate float64) float64 {
	return 100 * (((1 + investmentReturn) / (1 + inflationRate)) - 1)
}

// WACC - Weighted Average Cost of Capital
//
// Weighted Average Cost of Capital (WACC) is the rate that a company is expected to pay on average to all its security holders to finance its assets.
// e.g., If market value of equity is $600,000, market value of debt is $400,000, cost of equity is 6%, cost of debt is 5%, and tax rate is 35%, WACC is 4.9%.
func WACC(marketValueOfEquity, marketValueOfDebt, costOfEquity, costOfDebt, taxRate float64) float64 {
	E := marketValueOfEquity
	D := marketValueOfDebt
	V := marketValueOfEquity + marketValueOfDebt
	Re := costOfEquity
	Rd := costOfDebt
	T := taxRate
	wacc := ((E / V) * Re / 100) + (((D / V) * Rd / 100) * (1 - T/100))
	return math.Round(wacc*1000) / 10
}

// PI - Profitability Index
//
// Profitability Index (PI) is an index that attempts to identify the relationship between the costs and benefits of a proposed project through the use of a ratio calculated.
// e.g., If rate is 10%, initial investment is -$40,000, cash flows are $18,000, $12,000, $10,000, $9,000, and $6,000, PI is 1.09.
func PI(rate, initialInvestment float64, cashFlows []float64) float64 {
	var totalOfPVs float64 = 0

	for i := 0; i < len(cashFlows); i++ {
		index := float64(i)
		discountFactor := 1 / math.Pow(1+rate/100, index+1)
		totalOfPVs += cashFlows[i] * discountFactor
	}

	pi := totalOfPVs / math.Abs(initialInvestment)

	return math.Round(pi*100) / 100
}

// CAPM calculates expected return of an asset.
func CAPM(rf, beta, emr float64) float64 {
	return rf/100 + beta*(emr/100-rf/100)
}

// StockPV returns the Value of stock with dividend growing at a
//constant growth rate to perpetuity.
func StockPV(g, ke, D0 float64) float64 {
	return math.Round((D0 * (1 + g/100)) / ((ke / 100) - (g / 100)))
}
