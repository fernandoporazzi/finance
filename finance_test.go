package finance

import (
	"testing"
)

func TestPV(t *testing.T) {
	t.Run("Compute PV", func(t *testing.T) {
		entries := []struct {
			rate     float64
			cashFlow float64
			period   float64
			want     float64
		}{
			{5, 100, 1, 95.24},
			{5, 100, 5, 78.35},
		}

		for _, entry := range entries {
			got := PV(entry.rate, entry.cashFlow, entry.period)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestFV(t *testing.T) {
	t.Run("Compute FV", func(t *testing.T) {
		entries := []struct {
			rate     float64
			cashFlow float64
			period   float64
			want     float64
		}{
			{0.5, 1000, 12, 1061.68},
		}

		for _, entry := range entries {
			got := FV(entry.rate, entry.cashFlow, entry.period)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestNPV(t *testing.T) {
	t.Run("Compute NPV", func(t *testing.T) {
		entries := []struct {
			rate              float64
			initialInvestment float64
			cashFlows         []float64
			want              float64
		}{
			{10, -500000, []float64{200000, 300000, 200000}, 80015.03},
		}

		for _, entry := range entries {
			got := NPV(entry.rate, entry.initialInvestment, entry.cashFlows)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestROI(t *testing.T) {
	t.Run("Compute ROI", func(t *testing.T) {
		entries := []struct {
			initialInvestment float64
			earnings          float64
			want              float64
		}{
			{-55000, 60000, 9.09},
		}

		for _, entry := range entries {
			got := ROI(entry.initialInvestment, entry.earnings)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestPP(t *testing.T) {
	t.Run("Compute PP", func(t *testing.T) {
		entries := []struct {
			numberOfPeriods   int
			initialInvestment float64
			cashFlows         []float64
			want              float64
		}{
			{0, -105, []float64{25}, 4.2},                               // even cash flow
			{5, -50, []float64{10, 13, 16, 19, 22}, 3.4210526315789473}, // uneven cash flow
		}

		for _, entry := range entries {
			got := PP(entry.numberOfPeriods, entry.initialInvestment, entry.cashFlows)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestAM(t *testing.T) {
	t.Run("Compute AM", func(t *testing.T) {
		entries := []struct {
			principal      float64
			rate           float64
			period         float64
			paymentType    PaymentType
			payAtBeginning bool
			want           float64
		}{
			{20000, 7.5, 5, Years, false, 400.76},   // for inputs in years
			{20000, 7.5, 60, Months, false, 400.76}, // for inputs in months
			{20000, 7.5, 5, Years, true, 398.27},    // for inputs in years and payment at beginning
			{20000, 7.5, 60, Months, true, 398.27},  // for inputs in months and payment at beginning
		}

		for _, entry := range entries {
			got := AM(entry.principal, entry.rate, entry.period, entry.paymentType, entry.payAtBeginning)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestDF(t *testing.T) {
	t.Run("Compute DF", func(t *testing.T) {
		entries := []struct {
			rate            float64
			numberOfPeriods int
			want            []float64
		}{
			{10, 6, []float64{1, 0.91, 0.827, 0.752, 0.684}},
		}

		for _, entry := range entries {
			got := DF(entry.rate, entry.numberOfPeriods)

			for i, v := range got {
				if v != entry.want[i] {
					t.Errorf("Expected %v to be equal %v", v, entry.want[i])
				}
			}
		}
	})
}

func TestCI(t *testing.T) {
	t.Run("Compute CI", func(t *testing.T) {
		entries := []struct {
			rate, numOfCompoundings, principal, numOfPeriods float64
			want                                             float64
		}{
			{4.3, 4, 1500, 6, 1938.84},
		}

		for _, entry := range entries {
			got := CI(entry.rate, entry.numOfCompoundings, entry.principal, entry.numOfPeriods)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestCAGR(t *testing.T) {
	t.Run("Compute CAGR", func(t *testing.T) {
		entries := []struct {
			beginningValue, endingValue, numOfPeriods float64
			want                                      float64
		}{
			{10000, 19500, 3, 24.93},
		}

		for _, entry := range entries {
			got := CAGR(entry.beginningValue, entry.endingValue, entry.numOfPeriods)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestLR(t *testing.T) {
	t.Run("Compute LR", func(t *testing.T) {
		entries := []struct {
			totalLiabilities, totalDebts, totalIncome float64
			want                                      float64
		}{
			{25, 10, 20, 1.75},
		}

		for _, entry := range entries {
			got := LR(entry.totalLiabilities, entry.totalDebts, entry.totalIncome)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestR72(t *testing.T) {
	t.Run("Compute R72", func(t *testing.T) {
		entries := []struct {
			rate float64
			want float64
		}{
			{10, 7.2},
			{7, 10.285714285714286},
		}

		for _, entry := range entries {
			got := R72(entry.rate)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestPMT(t *testing.T) {
	t.Run("Compute PMT", func(t *testing.T) {
		entries := []struct {
			rate, numOfPayments, principal float64
			want                           float64
		}{
			{2, 36, -1000000, 39232.85},
		}

		for _, entry := range entries {
			got := PMT(entry.rate, entry.numOfPayments, entry.principal)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestIAR(t *testing.T) {
	t.Run("Compute IAR", func(t *testing.T) {
		entries := []struct {
			investmentReturn, inflationRate float64
			want                            float64
		}{
			{0.08, 0.03, 4.854368932038833},
		}

		for _, entry := range entries {
			got := IAR(entry.investmentReturn, entry.inflationRate)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestWACC(t *testing.T) {
	t.Run("Compute WACC", func(t *testing.T) {
		entries := []struct {
			marketValueOfEquity, marketValueOfDebt, costOfEquity, costOfDebt, taxRate float64
			want                                                                      float64
		}{
			{600000, 400000, 6, 5, 35, 4.9},
		}

		for _, entry := range entries {
			got := WACC(entry.marketValueOfEquity, entry.marketValueOfDebt, entry.costOfEquity, entry.costOfDebt, entry.taxRate)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestPI(t *testing.T) {
	t.Run("Compute PI", func(t *testing.T) {
		entries := []struct {
			rate, initialInvestment float64
			cashFlows               []float64
			want                    float64
		}{
			{10, -40000, []float64{18000, 12000, 10000, 9000, 6000}, 1.09},
		}

		for _, entry := range entries {
			got := PI(entry.rate, entry.initialInvestment, entry.cashFlows)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestCAPM(t *testing.T) {
	t.Run("Compute CAPM", func(t *testing.T) {
		entries := []struct {
			rf, beta, emr float64
			want          float64
		}{
			{2, 2, 10, 0.18},
		}

		for _, entry := range entries {
			got := CAPM(entry.rf, entry.beta, entry.emr)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}

func TestStockPV(t *testing.T) {
	t.Run("Compute StockPV", func(t *testing.T) {
		entries := []struct {
			g, ke, D0 float64
			want      float64
		}{
			{5, 15, 10, 105},
		}

		for _, entry := range entries {
			got := StockPV(entry.g, entry.ke, entry.D0)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}
