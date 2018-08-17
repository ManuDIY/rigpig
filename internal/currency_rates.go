package internal

type CurrencyRates struct {
	Rates []Currency
}

type Currency struct {
	name   string
	amount float64
}

var LatestCurrencyRates []Currency
