# finance
Common financial calculations powered by Go without dependencies.

![build](https://github.com/fernandoporazzi/finance/workflows/build/badge.svg)

## Supported calculations

- PV - Present Value
- FV - Future Value
- NPV - Net Present Value
- ROI - Return On Investment
- PP - Payback Period
- AM - Amortization
- DF - Discount Factor
- CI - Compound Interest
- CAGR - Compound Annual Growth Rate
- LR - Leverage Ratio
- R72 - Rule of 72
- PMT - Payment
- IAR - calculates the Inflation-adjusted return
- WACC - Weighted Average Cost of Capital
- PI - Profitability Index
- CAPM - calculates expected return of an asset
- StockPV - returns the Value of stock with dividend growing at a constant growth rate to perpetuity.

## Install

```
$ go get -u github.com/fernandoporazzi/finance
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/fernandoporazzi/finance"
)

func main() {
	fmt.Println(finance.ROI(-55000, 60000))
	fmt.Println(finance.AM(20000, 7.5, 6, finance.Months, false))
	cashFlows = []float64{18000, 12000, 10000, 9000, 6000}
	fmt.Println(finance.PI(10, -40000, cashFlows))
}
```

## Testing

```
$ go test -v . 
```

## License

[MIT License](https://github.com/fernandoporazzi/finance/blob/master/LICENSE)
