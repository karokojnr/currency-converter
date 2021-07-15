package main

import (
	"github.com/karokojnr/go-currency-converter/app/utils"
	"testing"
)

func TestConversion(t *testing.T) {
	conversionTables := []struct {
		amount float64
		fromCurrency string
		toCurrency string
		currencyRate float64
	}{
		{14, "KSH", "NGN", 53.34},
		{154, "GHS", "KSH", 2807.42},
		{201, "NGN", "GHS", 2.81},
	}

	for _, conversionTable := range conversionTables {
		result := utils.Convert(conversionTable.amount, conversionTable.fromCurrency, conversionTable.toCurrency)
		if result != conversionTable.currencyRate {
			t.Errorf("Conversion of %.2f  %v to %v  was incorrect, got: %.2f, expected: %.2f.", conversionTable.amount, conversionTable.fromCurrency, conversionTable.toCurrency, result, conversionTable.currencyRate)
		}
	}
}